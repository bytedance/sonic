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

#ifndef GO_FUNC_PARSER_H
#define GO_FUNC_PARSER_H

#include <map>
#include <string>
#include <vector>

namespace tool {

/**
 * @brief 表示一个寄存器
 *
 * 用于存储寄存器名称，支持基本的比较操作
 */
struct Register {
  std::string Name; ///< 寄存器名称（如 "x0", "d1" 等）

  /**
   * @brief 默认构造函数
   */
  Register() = default;

  /**
   * @brief 构造函数
   * @param Name 寄存器名称
   */
  Register(const std::string &Name) : Name(Name) {}

  /**
   * @brief 相等比较运算符
   * @param Other 另一个 Register 对象
   * @return 是否相等
   */
  bool operator==(const Register &Other) const { return Name == Other.Name; }

  /**
   * @brief 不等比较运算符
   * @param Other 另一个 Register 对象
   * @return 是否不等
   */
  bool operator!=(const Register &Other) const { return !(*this == Other); }
};

/**
 * @brief 表示一个参数或返回值（带名字和类型）
 *
 * 用于存储函数参数或返回值的信息，包括名称、类型、大小和寄存器分配
 */
struct Param {
  std::string Name;     ///< 参数名，若匿名则为空（如 "_" 或仅类型）
  std::string Type;     ///< 类型
  size_t Size = 0;      ///< 对齐后的大小（字节），0 表示未分配
  Register CReg;        ///< C ABI 使用的寄存器（x0-x7, d0-d7）
  bool IsFloat = false; ///< 是否为浮点类型（用于寄存器选择）

  /**
   * @brief 检查是否已分配寄存器
   * @return 是否已分配寄存器
   */
  bool HasRegister() const { return !CReg.Name.empty(); }
};

/**
 * @brief 函数签名信息
 *
 * 用于存储函数的签名信息，包括函数名、参数列表和返回值列表
 */
struct FuncSignature {
  std::string Name;           ///< 函数名
  std::vector<Param> Params;  ///< 参数列表
  std::vector<Param> Results; ///< 返回值列表（支持命名）

  /**
   * @brief 计算总参数+返回值空间（字节）
   * @return 总空间大小
   */
  size_t ArgSpace() const {
    size_t Total = InputSpace();
    for (const auto &R : Results) {
      Total += R.Size;
    }
    return Total;
  }

  /**
   * @brief 计算仅参数空间（字节）
   * @return 参数空间大小
   */
  size_t InputSpace() const {
    size_t Total = 0;
    for (const auto &P : Params) {
      Total += P.Size;
    }
    return Total;
  }

  /**
   * @brief 检查是否所有参数/返回值都已分配寄存器
   * @return 是否分配成功
   */
  bool IsAllocated() const {
    for (const auto &P : Params) {
      if (!P.HasRegister()) {
        return false;
      }
    }
    for (const auto &R : Results) {
      if (!R.HasRegister()) {
        return false;
      }
    }
    return true;
  }
};

/**
 * @brief 解析结果
 *
 * 用于存储解析 Go 函数签名的结果，成功时 funcs 非空，失败时 error 非空
 */
struct ParseResult {
  std::map<std::string, FuncSignature> Funcs; ///< 函数签名映射
  std::string Error;                          ///< 若非空，表示解析失败原因

  /**
   * @brief 检查解析是否成功
   * @return 是否成功
   */
  bool Success() const { return Error.empty(); }
};

/**
 * @brief 基于 ParseResult 进行寄存器分配
 *
 * 为解析结果中的所有函数参数和返回值分配寄存器，失败时填充 Error
 * @param Result 解析结果
 */
void AllocateRegisters(ParseResult &Result) noexcept;

/**
 * @brief 从字符串内容解析 Go 汇编绑定函数（无函数体）
 *
 * 解析 Go 源码中的函数签名，只处理没有函数体的声明
 * @param Content Go 源码内容（非空）
 * @return 解析结果，包含函数映射或错误信息
 */
ParseResult ParseGoAsmFunctions(const std::string &Content) noexcept;

/**
 * @brief 从文件路径读取并解析 Go 函数签名
 *
 * 读取指定路径的 Go 文件，并解析其中的函数签名
 * @param FilePath Go 文件路径
 * @return 解析结果
 */
ParseResult ParseGoFile(const std::string &FilePath) noexcept;

} // end namespace tool

#endif // GO_FUNC_PARSER_H