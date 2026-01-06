from tokenizer import TokenType, Token
import std/[strutils, options]

type Visitor = object
    tokens: seq[Token]
    index: int = 0

proc peek(visitor: var Visitor): Token = visitor.tokens[visitor.index]

proc consume(visitor: var Visitor): Token =
    result = visitor.tokens[visitor.index]
    inc visitor.index

type ParseResult*[T] = object
    remainingTokens*: seq[Token]
    node*: T

type IntLiteralNode* = object
    value*: int

type ExitNode* = object
    exit_code*: IntLiteralNode

type IdentNode* = object
    identifier*: string

type TypeNode* = enum
    INT,
    VOID

type TypeDeclNode* = object
    typedecl: TypeNode

type FuncNode* = object
    return_type*: TypeDeclNode
    ident*: IdentNode

proc parseExit(visitor: Visitor): Option[ParseResult[ExitNode]] =
    return none(ParseResult[ExitNode])

proc parse* (tokens: seq[Token]): ExitNode {. raises: [Exception] .} =
    raise newException(Exception, "Parsing failed!")