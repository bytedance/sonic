/*
 * Copyright 2026 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "go_func_parser.h"

#include "llvm/Support/raw_ostream.h"

#include <cctype>
#include <fstream>
#include <sstream>
#include <unordered_map>

namespace tool {

/**
 * @brief 去除字符串首尾的空白字符
 *
 * @param Str 输入字符串
 * @return 去除空白后的字符串
 */
std::string Trim(const std::string &Str) {
  if (Str.empty()) {
    return Str;
  }
  auto Start = Str.begin();
  while (Start != Str.end() &&
         std::isspace(static_cast<unsigned char>(*Start))) {
    ++Start;
  }
  auto End = Str.end();
  do {
    --End;
  } while (std::distance(Start, End) > 0 &&
           std::isspace(static_cast<unsigned char>(*End)));
  return Start <= End ? std::string(Start, End + 1) : std::string();
}

/**
 * @brief 判断是否为编译器指令或 IDE 注解
 *
 * @param Line 输入行
 * @return 是否为编译器指令
 */
bool IsCompilerDirective(const std::string &Line) {
  std::string Trimmed = Trim(Line);
  return Trimmed.rfind("//go:", 0) == 0 || Trimmed.rfind("//goland:", 0) == 0;
}

/**
 * @brief 检查字符串是否包含有效 func 开头（忽略前置空格和指令）
 *
 * @param Line 输入行
 * @return 是否以 func 开头
 */
bool StartsWithFunc(const std::string &Line) {
  std::string Trimmed = Trim(Line);
  return Trimmed.size() >= 5 && Trimmed.substr(0, 5) == "func ";
}

/**
 * @brief 计算字符串中未闭合的圆括号数量
 *
 * 仅用于判断函数声明是否结束
 * @param Str 输入字符串
 * @return 未闭合的圆括号数量
 */
int CountUnmatchedParens(const std::string &Str) {
  int Paren = 0;
  for (char C : Str) {
    if (C == '(') {
      ++Paren;
    } else if (C == ')') {
      --Paren;
    }
  }
  return Paren;
}

/**
 * @brief 解析单个参数或返回值字段
 *
 * 从字段中分离出名称和类型
 * @param Field 输入字段
 * @return 解析后的 Param 对象
 */
Param ParseField(const std::string &Field) {
  if (Field.empty()) {
    return {"", ""};
  }

  int Paren = 0, Bracket = 0, Brace = 0;
  size_t LastSpace = std::string::npos;

  for (size_t I = Field.size(); I-- > 0;) {
    unsigned char C = static_cast<unsigned char>(Field[I]);
    if (C == ')') {
      ++Paren;
    } else if (C == '(') {
      --Paren;
    } else if (C == ']') {
      ++Bracket;
    } else if (C == '[') {
      --Bracket;
    } else if (C == '}') {
      ++Brace;
    } else if (C == '{') {
      --Brace;
    } else if (C == ' ' && Paren == 0 && Bracket == 0 && Brace == 0) {
      LastSpace = I;
      break;
    }
  }

  if (LastSpace == std::string::npos) {
    return {"", Field};
  } else {
    std::string NamePart = Trim(Field.substr(0, LastSpace));
    std::string TypePart = Trim(Field.substr(LastSpace + 1));
    return {NamePart, TypePart};
  }
}

/**
 * @brief 分割参数列表
 *
 * 按照逗号分割参数列表，考虑括号内的逗号
 * @param Str 输入字符串
 * @return 分割后的参数列表
 */
std::vector<std::string> SplitParams(const std::string &Str) {
  if (Str.empty()) {
    return {};
  }
  std::vector<std::string> Parts;
  int Paren = 0, Bracket = 0, Brace = 0;
  size_t Start = 0;
  for (size_t I = 0; I <= Str.size(); ++I) {
    char C = (I == Str.size()) ? ',' : Str[I];
    bool AtEnd = (I == Str.size());
    if (!AtEnd) {
      if (C == '(') {
        ++Paren;
      } else if (C == ')') {
        --Paren;
      } else if (C == '[') {
        ++Bracket;
      } else if (C == ']') {
        --Bracket;
      } else if (C == '{') {
        ++Brace;
      } else if (C == '}') {
        --Brace;
      }
    }
    if ((C == ',' && Paren == 0 && Bracket == 0 && Brace == 0) || AtEnd) {
      std::string Part = Trim(Str.substr(Start, I - Start));
      if (!Part.empty()) {
        Parts.push_back(Part);
      }
      Start = I + 1;
    }
  }
  return Parts;
}

/**
 * @brief 解析参数列表
 *
 * 将参数列表字符串解析为 Param 对象列表
 * @param List 参数列表字符串
 * @return 解析后的 Param 对象列表
 */
std::vector<Param> ParseParamList(const std::string &List) {
  auto Items = SplitParams(List);
  std::vector<Param> Params;
  for (const auto &Item : Items) {
    Params.push_back(ParseField(Item));
  }
  return Params;
}

/**
 * @brief 解析返回值列表
 *
 * 将返回值列表字符串解析为 Param 对象列表
 * @param ResultList 返回值列表字符串
 * @return 解析后的 Param 对象列表
 */
std::vector<Param> ParseResultList(const std::string &ResultList) {
  if (ResultList.empty()) {
    return {};
  }

  if (ResultList.front() != '(') {
    std::string Type = Trim(ResultList);
    return {{"", Type}};
  }

  if (ResultList.size() < 2 || ResultList.back() != ')') {
    return {{"", ResultList}};
  }

  std::string Inner = Trim(ResultList.substr(1, ResultList.size() - 2));
  if (Inner.empty()) {
    return {};
  }

  return ParseParamList(Inner);
}

ParseResult ParseGoAsmFunctions(const std::string &Content) noexcept {
  ParseResult Result;
  if (Content.empty()) {
    Result.Error = "Input content is empty";
    return Result;
  }

  std::istringstream Iss(Content);
  std::string Line;
  int LineNumber = 0;

  std::string CurrentFuncLine;
  bool InFunc = false;
  int FuncStartLine = 0;

  auto TryParseFunction = [&](const std::string &FullLine,
                              int StartLine) -> bool {
    std::string L = FullLine;
    // 移除行尾注释（简单版，跨行时可能不准确，但够用）
    size_t Comment = L.find("//");
    if (Comment != std::string::npos) {
      // 只有当 // 不在字符串或类型内部时才移除（简化处理）
      L = L.substr(0, Comment);
    }
    L = Trim(L);
    if (L.empty()) {
      return false;
    }

    if (!StartsWithFunc(L)) {
      return false;
    }

    size_t Pos = 5;
    while (Pos < L.size() && std::isspace(static_cast<unsigned char>(L[Pos]))) {
      ++Pos;
    }
    if (Pos >= L.size()) {
      return false;
    }

    size_t NameEnd = Pos;
    while (NameEnd < L.size() &&
           (std::isalnum(static_cast<unsigned char>(L[NameEnd])) ||
            L[NameEnd] == '_')) {
      ++NameEnd;
    }
    std::string Name = L.substr(Pos, NameEnd - Pos);
    if (Name.empty()) {
      return false;
    }

    size_t FirstParen = L.find('(', NameEnd);
    if (FirstParen == std::string::npos) {
      return false;
    }

    // 找到参数列表结束位置（匹配括号）
    int ParenCount = 1;
    size_t I = FirstParen + 1;
    while (I < L.size() && ParenCount > 0) {
      if (L[I] == '(') {
        ++ParenCount;
      } else if (L[I] == ')') {
        --ParenCount;
      }
      ++I;
    }
    if (ParenCount != 0) {
      return false; // not balanced
    }

    std::string ParamStr = L.substr(FirstParen, I - FirstParen);
    std::string Rest = Trim(L.substr(I));

    std::string ResultStr;
    if (!Rest.empty()) {
      if (Rest.front() == '(') {
        int Rp = 1;
        size_t J = 1;
        while (J < Rest.size() && Rp > 0) {
          if (Rest[J] == '(') {
            ++Rp;
          } else if (Rest[J] == ')') {
            --Rp;
          }
          ++J;
        }
        if (Rp == 0) {
          ResultStr = Rest.substr(0, J);
        } else {
          return false;
        }
      } else {
        ResultStr = Rest;
      }
    }

    // 检查是否有函数体：签名结束后是否有 {
    size_t SignatureEnd = I;
    if (!ResultStr.empty()) {
      size_t Rp = L.find(ResultStr, I);
      if (Rp != std::string::npos) {
        SignatureEnd = Rp + ResultStr.size();
      }
    }

    bool HasBody = false;
    for (size_t K = SignatureEnd; K < L.size(); ++K) {
      if (L[K] == '{') {
        HasBody = true;
        break;
      }
      if (!std::isspace(static_cast<unsigned char>(L[K]))) {
        break;
      }
    }
    if (HasBody) {
      return false;
    }

    // 解析
    std::string ParamInner = ParamStr.substr(1, ParamStr.size() - 2);
    auto Params =
        ParamInner.empty() ? std::vector<Param>{} : ParseParamList(ParamInner);
    auto Results = ParseResultList(ResultStr);

    if (Result.Funcs.count(Name)) {
      Result.Error = "Line " + std::to_string(StartLine) +
                     ": duplicate function '" + Name + "'";
      return false;
    }

    Result.Funcs.emplace(Name, FuncSignature{Name, Params, Results});
    return true;
  };

  while (std::getline(Iss, Line)) {
    ++LineNumber;

    // 跳过纯编译器指令行
    if (IsCompilerDirective(Line)) {
      continue;
    }

    std::string Trimmed = Trim(Line);
    if (Trimmed.empty()) {
      if (InFunc) {
        // 空行中断函数声明
        InFunc = false;
        CurrentFuncLine.clear();
      }
      continue;
    }

    if (InFunc) {
      CurrentFuncLine += " " + Line; // 保留原始内容（含注释）
      int Unmatched = CountUnmatchedParens(CurrentFuncLine);
      if (Unmatched == 0) {
        // 尝试解析
        if (TryParseFunction(CurrentFuncLine, FuncStartLine)) {
          // success
        }
        InFunc = false;
        CurrentFuncLine.clear();
      }
      // else: 继续收集下一行
    } else {
      // 不在函数中
      if (StartsWithFunc(Trimmed)) {
        CurrentFuncLine = Line;
        int Unmatched = CountUnmatchedParens(CurrentFuncLine);
        if (Unmatched == 0) {
          if (TryParseFunction(CurrentFuncLine, LineNumber)) {
            // success
          }
        } else {
          InFunc = true;
          FuncStartLine = LineNumber;
        }
      }
      // else: ignore
    }
  }

  // 处理文件末尾未闭合的函数（可选）
  // 这里选择忽略

  return Result;
}

ParseResult ParseGoFile(const std::string &FilePath) noexcept {
  if (FilePath.empty()) {
    return ParseResult{{}, "File path is empty"};
  }

  std::ifstream File(FilePath);
  if (!File.is_open()) {
    return ParseResult{{}, "Cannot open file: " + FilePath};
  }

  std::string Content((std::istreambuf_iterator<char>(File)),
                      std::istreambuf_iterator<char>());
  return ParseGoAsmFunctions(Content);
}

// 整数寄存器分配顺序
const std::vector<Register> INTEGER_REGISTER_ORDER = {
    {"x0"}, {"x1"}, {"x2"}, {"x3"}, {"x4"}, {"x5"}, {"x6"}, {"x7"}};

// 浮点寄存器分配顺序
const std::vector<Register> FLOAT_REGISTER_ORDER = {
    {"d0"}, {"d1"}, {"d2"}, {"d3"}, {"d4"}, {"d5"}, {"d6"}, {"d7"}};

// 类型信息映射：类型名 → (大小, 是否为浮点类型)
std::unordered_map<std::string, std::pair<size_t, bool>> TypeInfo = {
    {"int", {8, false}},
    {"uint", {8, false}},
    {"uintptr", {8, false}},
    {"Pointer", {8, false}},
    {"unsafe.Pointer", {8, false}},
    {"int64", {8, false}},
    {"uint64", {8, false}},
    {"int32", {4, false}},
    {"uint32", {4, false}},
    {"rune", {4, false}},
    {"int16", {2, false}},
    {"uint16", {2, false}},
    {"int8", {1, false}},
    {"uint8", {1, false}},
    {"byte", {1, false}},
    {"bool", {1, false}},
    {"float32", {4, true}},
    {"float64", {8, true}}};

/**
 * @brief 获取类型信息
 *
 * 根据类型名获取类型的大小和是否为浮点类型
 * @param TypeStr 类型名
 * @return 类型信息：(大小, 是否为浮点类型)
 */
static std::pair<size_t, bool> GetTypeInfo(const std::string &TypeStr) {
  if (TypeStr.empty()) {
    llvm::outs() << "empty type\n";
    return {};
  }

  std::string Type = TypeStr;
  // 处理指针
  if (Type[0] == '*') {
    return {8, false};
  }

  // 标准类型映射
  if (TypeInfo.find(Type) != TypeInfo.end()) {
    return TypeInfo[Type];
  } else {
    llvm::outs() << "unrecognized type: " << TypeStr << "\n";
    return {};
  }
}

/**
 * @brief 对齐大小到 8 字节
 *
 * 因为寄存器大小为 8 字节
 * @param Size 原始大小
 * @return 对齐后的大小
 */
static size_t AlignSize(size_t Size) { return ((Size + 7) / 8) * 8; }

/**
 * @brief 分配单个 Param 的寄存器
 *
 * 为参数或返回值分配适当的寄存器
 * @param Param 参数或返回值
 * @param IntIdx 整数寄存器索引
 * @param FpIdx 浮点寄存器索引
 * @param FuncName 函数名
 */
static void AllocateParam(Param &Param, int &IntIdx, int &FpIdx,
                          const std::string &FuncName) {
  auto [RawSize, IsFloat] = GetTypeInfo(Param.Type);
  if (RawSize == 0) {
    return;
  }
  Param.IsFloat = IsFloat;
  Param.Size = AlignSize(RawSize);

  if (IsFloat) {
    if (FpIdx >= static_cast<int>(FLOAT_REGISTER_ORDER.size())) {
      llvm::outs() << "too many floating-point arguments\n";
      return;
    }
    Param.CReg = FLOAT_REGISTER_ORDER[FpIdx];
    ++FpIdx;
  } else {
    if (IntIdx >= static_cast<int>(INTEGER_REGISTER_ORDER.size())) {
      llvm::outs() << "too many integer arguments\n";
      return;
    }
    Param.CReg = INTEGER_REGISTER_ORDER[IntIdx];
    ++IntIdx;
  }
}

void AllocateRegisters(ParseResult &Result) noexcept {
  if (!Result.Success()) {
    return;
  }

  for (auto &[Name, Sig] : Result.Funcs) {
    // 检查返回值数量（只支持 ≤1）
    if (Sig.Results.size() > 1) {
      Result.Error =
          "Function '" + Name +
          "' has multiple return values (only single return supported)";
      return;
    }

    // 分配参数寄存器
    int IntIdx = 0, FpIdx = 0;
    for (auto &Param : Sig.Params) {
      AllocateParam(Param, IntIdx, FpIdx, Name);
    }

    // 分配返回值寄存器（如果有）
    if (!Sig.Results.empty()) {
      // 重置索引（返回值从第 0 个寄存器开始）
      int RetIntIdx = 0, RetFpIdx = 0;
      AllocateParam(Sig.Results[0], RetIntIdx, RetFpIdx, Name);
    }
  }
}

} // end namespace tool