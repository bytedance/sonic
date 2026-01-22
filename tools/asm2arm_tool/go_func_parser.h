#ifndef GO_FUNC_PARSER_H
#define GO_FUNC_PARSER_H

#include <string>
#include <vector>
#include <map>

struct Register {
    std::string name;
    Register() = default;
    Register(const std::string& n) : name(n)
    {}
    bool operator==(const Register& other) const
    {
        return name == other.name;
    }
    bool operator!=(const Register& other) const
    {
        return !(*this == other);
    }
};

/// 表示一个参数或返回值（带名字和类型）
struct Param {
    std::string name;       ///< 参数名，若匿名则为空（如 "_" 或仅类型）
    std::string type;       ///< 类型
    size_t size = 0;        ///< 对齐后的大小（字节），0 表示未分配
    Register creg;          ///< C ABI 使用的寄存器（x0-x7, d0-d7）
    bool is_float = false;  ///< 是否为浮点类型（用于寄存器选择）

    bool hasRegister() const
    {
        return !creg.name.empty();
    }
};

/// 函数签名信息
struct FuncSignature {
    std::string name;            ///< 函数名
    std::vector<Param> params;   ///< 参数列表
    std::vector<Param> results;  ///< 返回值列表（支持命名）

    /// 总参数+返回值空间（字节）
    size_t argspace() const
    {
        size_t total = inputspace();
        for (const auto& r : results) {
            total += r.size;
        }
        return total;
    }

    /// 仅参数空间（字节）
    size_t inputspace() const
    {
        size_t total = 0;
        for (const auto& p : params) {
            total += p.size;
        }
        return total;
    }

    /// 是否分配成功（所有参数/返回值都有寄存器）
    bool isAllocated() const
    {
        for (const auto& p : params) {
            if (!p.hasRegister()) {
                return false;
            }
        }
        for (const auto& r : results) {
            if (!r.hasRegister()) {
                return false;
            }
        }
        return true;
    }
};

/// 解析结果：成功时 funcs 非空，失败时 error 非空
struct ParseResult {
    std::map<std::string, FuncSignature> funcs;
    std::string error;  ///< 若非空，表示解析失败原因

    bool success() const
    {
        return error.empty();
    }
};

/// 基于 ParseResult 进行寄存器分配，失败时填充 error
void allocateRegisters(ParseResult& result) noexcept;

/// 从字符串内容解析 Go 汇编绑定函数（无函数体）
/// @param content Go 源码内容（非空）
/// @return 解析结果，包含函数映射或错误信息
ParseResult parseGoAsmFunctions(const std::string& content) noexcept;

/// 从文件路径读取并解析
/// @param filepath Go 文件路径
/// @return 解析结果
ParseResult parseGoFile(const std::string& filepath) noexcept;

#endif  // GO_FUNC_PARSER_H