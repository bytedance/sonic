#ifndef GO_FUNC_PARSER_H
#define GO_FUNC_PARSER_H

#include <string>
#include <vector>
#include <map>

/// 表示一个参数或返回值（带名字和类型）
struct Param {
    std::string name;  ///< 参数名，若匿名则为空（如 "_" 或仅类型）
    std::string type;  ///< 类型
};

/// 函数签名信息
struct FuncSignature {
    std::string name;            ///< 函数名
    std::vector<Param> params;   ///< 参数列表
    std::vector<Param> results;  ///< 返回值列表（支持命名）
};

/// 解析结果：成功时 funcs 非空，失败时 error 非空
struct ParseResult {
    std::map<std::string, FuncSignature> funcs;
    std::string error;  ///< 若非空，表示解析失败原因
};

/// 从字符串内容解析 Go 汇编绑定函数（无函数体）
/// @param content Go 源码内容（非空）
/// @return 解析结果，包含函数映射或错误信息
ParseResult parseGoAsmFunctions(const std::string& content) noexcept;

/// 从文件路径读取并解析
/// @param filepath Go 文件路径
/// @return 解析结果
ParseResult parseGoFile(const std::string& filepath) noexcept;

#endif  // GO_FUNC_PARSER_H