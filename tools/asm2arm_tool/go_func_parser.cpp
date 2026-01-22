#include "go_func_parser.h"

#include "llvm/Support/raw_ostream.h"

#include <fstream>
#include <sstream>
#include <cctype>
#include <unordered_map>

std::string trim(const std::string& s)
{
    if (s.empty()) {
        return s;
    }
    auto start = s.begin();
    while (start != s.end() && std::isspace(static_cast<unsigned char>(*start))) {
        ++start;
    }
    auto end = s.end();
    do {
        --end;
    } while (std::distance(start, end) > 0 && std::isspace(static_cast<unsigned char>(*end)));
    return start <= end ? std::string(start, end + 1) : std::string();
}

// 判断是否为编译器指令或 IDE 注解
bool isCompilerDirective(const std::string& line)
{
    std::string t = trim(line);
    return t.rfind("//go:", 0) == 0 || t.rfind("//goland:", 0) == 0;
}

// 检查字符串是否包含有效 func 开头（忽略前置空格和指令）
bool startsWithFunc(const std::string& line)
{
    std::string t = trim(line);
    return t.size() >= 5 && t.substr(0, 5) == "func ";
}

// 括号匹配：返回 true 如果括号平衡
bool isBalancedParens(const std::string& s)
{
    int paren = 0, bracket = 0, brace = 0;
    for (char c : s) {
        switch (c) {
            case '(':
                ++paren;
                break;
            case ')':
                --paren;
                if (paren < 0)
                    return false;
                break;
            case '[':
                ++bracket;
                break;
            case ']':
                --bracket;
                if (bracket < 0)
                    return false;
                break;
            case '{':
                ++brace;
                break;
            case '}':
                --brace;
                if (brace < 0)
                    return false;
                break;
        }
    }
    return paren == 0 && bracket == 0 && brace == 0;
}

// 计算字符串中未闭合的圆括号数量（仅用于判断是否结束）
int countUnmatchedParens(const std::string& s)
{
    int paren = 0;
    for (char c : s) {
        if (c == '(') {
            ++paren;
        } else if (c == ')') {
            --paren;
        }
    }
    return paren;
}

Param parseField(const std::string& field)
{
    if (field.empty()) {
        return {"", ""};
    }

    int paren = 0, bracket = 0, brace = 0;
    size_t lastSpace = std::string::npos;

    for (size_t i = field.size(); i-- > 0;) {
        unsigned char c = static_cast<unsigned char>(field[i]);
        if (c == ')') {
            ++paren;
        } else if (c == '(') {
            --paren;
        } else if (c == ']') {
            ++bracket;
        } else if (c == '[') {
            --bracket;
        } else if (c == '}') {
            ++brace;
        } else if (c == '{') {
            --brace;
        } else if (c == ' ' && paren == 0 && bracket == 0 && brace == 0) {
            lastSpace = i;
            break;
        }
    }

    if (lastSpace == std::string::npos) {
        return {"", field};
    } else {
        std::string namePart = trim(field.substr(0, lastSpace));
        std::string typePart = trim(field.substr(lastSpace + 1));
        return {namePart, typePart};
    }
}

std::vector<std::string> splitParams(const std::string& s)
{
    if (s.empty()) {
        return {};
    }
    std::vector<std::string> parts;
    int paren = 0, bracket = 0, brace = 0;
    size_t start = 0;
    for (size_t i = 0; i <= s.size(); ++i) {
        char c = (i == s.size()) ? ',' : s[i];
        bool atEnd = (i == s.size());
        if (!atEnd) {
            if (c == '(') {
                ++paren;
            } else if (c == ')') {
                --paren;
            } else if (c == '[') {
                ++bracket;
            } else if (c == ']') {
                --bracket;
            } else if (c == '{') {
                ++brace;
            } else if (c == '}') {
                --brace;
            }
        }
        if ((c == ',' && paren == 0 && bracket == 0 && brace == 0) || atEnd) {
            std::string part = trim(s.substr(start, i - start));
            if (!part.empty()) {
                parts.push_back(part);
            }
            start = i + 1;
        }
    }
    return parts;
}

std::vector<Param> parseParamList(const std::string& list)
{
    auto items = splitParams(list);
    std::vector<Param> params;
    for (const auto& item : items) {
        params.push_back(parseField(item));
    }
    return params;
}

std::vector<Param> parseResultList(const std::string& resultList)
{
    if (resultList.empty()) {
        return {};
    }

    if (resultList.front() != '(') {
        std::string type = trim(resultList);
        return {{"", type}};
    }

    if (resultList.size() < 2 || resultList.back() != ')') {
        return {{"", resultList}};
    }

    std::string inner = trim(resultList.substr(1, resultList.size() - 2));
    if (inner.empty()) {
        return {};
    }

    return parseParamList(inner);
}

ParseResult parseGoAsmFunctions(const std::string& content) noexcept
{
    ParseResult result;
    if (content.empty()) {
        result.error = "Input content is empty";
        return result;
    }

    std::istringstream iss(content);
    std::string line;
    int lineNumber = 0;

    std::string currentFuncLine;
    bool inFunc = false;
    int funcStartLine = 0;

    auto tryParseFunction = [&](const std::string& fullLine, int startLine) -> bool {
        std::string l = fullLine;
        // 移除行尾注释（简单版，跨行时可能不准确，但够用）
        size_t comment = l.find("//");
        if (comment != std::string::npos) {
            // 只有当 // 不在字符串或类型内部时才移除（简化处理）
            l = l.substr(0, comment);
        }
        l = trim(l);
        if (l.empty()) {
            return false;
        }

        if (!startsWithFunc(l)) {
            return false;
        }

        size_t pos = 5;
        while (pos < l.size() && std::isspace(static_cast<unsigned char>(l[pos]))) {
            ++pos;
        }
        if (pos >= l.size()) {
            return false;
        }

        size_t nameEnd = pos;
        while (nameEnd < l.size() && (std::isalnum(static_cast<unsigned char>(l[nameEnd])) || l[nameEnd] == '_')) {
            ++nameEnd;
        }
        std::string name = l.substr(pos, nameEnd - pos);
        if (name.empty()) {
            return false;
        }

        size_t firstParen = l.find('(', nameEnd);
        if (firstParen == std::string::npos) {
            return false;
        }

        // 找到参数列表结束位置（匹配括号）
        int parenCount = 1;
        size_t i = firstParen + 1;
        while (i < l.size() && parenCount > 0) {
            if (l[i] == '(') {
                ++parenCount;
            } else if (l[i] == ')') {
                --parenCount;
            }
            ++i;
        }
        if (parenCount != 0) {
            return false;  // not balanced
        }

        std::string paramStr = l.substr(firstParen, i - firstParen);
        std::string rest = trim(l.substr(i));

        std::string resultStr;
        if (!rest.empty()) {
            if (rest.front() == '(') {
                int rp = 1;
                size_t j = 1;
                while (j < rest.size() && rp > 0) {
                    if (rest[j] == '(') {
                        ++rp;
                    } else if (rest[j] == ')') {
                        --rp;
                    }
                    ++j;
                }
                if (rp == 0) {
                    resultStr = rest.substr(0, j);
                } else {
                    return false;
                }
            } else {
                resultStr = rest;
            }
        }

        // 检查是否有函数体：签名结束后是否有 {
        size_t signatureEnd = i;
        if (!resultStr.empty()) {
            size_t rp = l.find(resultStr, i);
            if (rp != std::string::npos) {
                signatureEnd = rp + resultStr.size();
            }
        }

        bool hasBody = false;
        for (size_t k = signatureEnd; k < l.size(); ++k) {
            if (l[k] == '{') {
                hasBody = true;
                break;
            }
            if (!std::isspace(static_cast<unsigned char>(l[k]))) {
                break;
            }
        }
        if (hasBody) {
            return false;
        }

        // 解析
        std::string paramInner = paramStr.substr(1, paramStr.size() - 2);
        auto params = paramInner.empty() ? std::vector<Param>{} : parseParamList(paramInner);
        auto results = parseResultList(resultStr);

        if (result.funcs.count(name)) {
            result.error = "Line " + std::to_string(startLine) + ": duplicate function '" + name + "'";
            return false;
        }

        result.funcs.emplace(name, FuncSignature{name, params, results});
        return true;
    };

    while (std::getline(iss, line)) {
        ++lineNumber;

        // 跳过纯编译器指令行
        if (isCompilerDirective(line)) {
            continue;
        }

        std::string trimmed = trim(line);
        if (trimmed.empty()) {
            if (inFunc) {
                // 空行中断函数声明
                inFunc = false;
                currentFuncLine.clear();
            }
            continue;
        }

        if (inFunc) {
            currentFuncLine += " " + line;  // 保留原始内容（含注释）
            int unmatched = countUnmatchedParens(currentFuncLine);
            if (unmatched == 0) {
                // 尝试解析
                if (tryParseFunction(currentFuncLine, funcStartLine)) {
                    // success
                }
                inFunc = false;
                currentFuncLine.clear();
            }
            // else: 继续收集下一行
        } else {
            // 不在函数中
            if (startsWithFunc(trimmed)) {
                currentFuncLine = line;
                int unmatched = countUnmatchedParens(currentFuncLine);
                if (unmatched == 0) {
                    if (tryParseFunction(currentFuncLine, lineNumber)) {
                        // success
                    }
                } else {
                    inFunc = true;
                    funcStartLine = lineNumber;
                }
            }
            // else: ignore
        }
    }

    // 处理文件末尾未闭合的函数（可选）
    // 这里选择忽略

    return result;
}

ParseResult parseGoFile(const std::string& filepath) noexcept
{
    if (filepath.empty()) {
        return ParseResult{{}, "File path is empty"};
    }

    std::ifstream file(filepath);
    if (!file.is_open()) {
        return ParseResult{{}, "Cannot open file: " + filepath};
    }

    std::string content((std::istreambuf_iterator<char>(file)), std::istreambuf_iterator<char>());
    return parseGoAsmFunctions(content);
}

const std::vector<Register> ARGS_ORDER_C = {{"x0"}, {"x1"}, {"x2"}, {"x3"}, {"x4"}, {"x5"}, {"x6"}, {"x7"}};
const std::vector<Register> FPARGS_ORDER = {{"d0"}, {"d1"}, {"d2"}, {"d3"}, {"d4"}, {"d5"}, {"d6"}, {"d7"}};

std::unordered_map<std::string, std::pair<size_t, bool>> TypeInfo = {{"int", {8, false}}, {"uint", {8, false}},
    {"uintptr", {8, false}}, {"Pointer", {8, false}}, {"unsafe.Pointer", {8, false}}, {"int64", {8, false}},
    {"uint64", {8, false}}, {"int32", {4, false}}, {"uint32", {4, false}}, {"rune", {4, false}}, {"int16", {2, false}},
    {"uint16", {2, false}}, {"int8", {1, false}}, {"uint8", {1, false}}, {"byte", {1, false}}, {"bool", {1, false}},
    {"float32", {4, true}}, {"float64", {8, true}}};

// 类型 → (size, isFloat)
static std::pair<size_t, bool> getTypeInfo(const std::string& typeStr)
{
    if (typeStr.empty()) {
        llvm::outs() << "empty type\n";
        return {};
    }

    std::string t = typeStr;
    // 处理指针
    if (t[0] == '*') {
        return {8, false};
    }

    // 标准类型映射
    if (TypeInfo.find(t) != TypeInfo.end()) {
        return TypeInfo[t];
    } else {
        llvm::outs() << "unrecognized type: " << typeStr << "\n";
        return {};
    }
}

// 对齐到 8 字节（寄存器大小）
static size_t alignSize(size_t size)
{
    return ((size + 7) / 8) * 8;
}

// 分配单个 Param 的寄存器
static void allocateParam(Param& param, int& intIdx, int& fpIdx, const std::string& funcName)
{
    auto [rawSize, isFloat] = getTypeInfo(param.type);
    if (rawSize == 0) {
        return;
    }
    param.is_float = isFloat;
    param.size = alignSize(rawSize);

    if (isFloat) {
        if (fpIdx >= static_cast<int>(FPARGS_ORDER.size())) {
            llvm::outs() << "too many floating-point arguments\n";
            return;
        }
        param.creg = FPARGS_ORDER[fpIdx];
        ++fpIdx;
    } else {
        if (intIdx >= static_cast<int>(ARGS_ORDER_C.size())) {
            llvm::outs() << "too many integer arguments\n";
            return;
        }
        param.creg = ARGS_ORDER_C[intIdx];
        ++intIdx;
    }
}

// 主分配函数
void allocateRegisters(ParseResult& result) noexcept
{
    if (!result.success()) {
        return;
    }

    for (auto& [name, sig] : result.funcs) {
        // 检查返回值数量（只支持 ≤1）
        if (sig.results.size() > 1) {
            result.error = "Function '" + name + "' has multiple return values (only single return supported)";
            return;
        }

        // 分配参数寄存器
        int intIdx = 0, fpIdx = 0;
        for (auto& param : sig.params) {
            allocateParam(param, intIdx, fpIdx, name);
        }

        // 分配返回值寄存器（如果有）
        if (!sig.results.empty()) {
            // 重置索引（返回值从第 0 个寄存器开始）
            int retIntIdx = 0, retFpIdx = 0;
            allocateParam(sig.results[0], retIntIdx, retFpIdx, name);
        }
    }
}