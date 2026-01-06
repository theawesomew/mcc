import std/options
import std/strutils

type TokenType* = enum
  EXIT,
  INT_LITERAL,
  SEMICOLON,
  INT_TYPE,
  VOID_TYPE,
  IDENT,
  OPEN_CURLY,
  CLOSE_CURLY,
  OPEN_PAREN,
  CLOSE_PAREN


type Token* = object
  tokenType*: TokenType
  value*: Option[string]

proc tokenize* (input: string): seq[Token] {. raises: [Exception] .} =
  result = @[]

  var i = 0
  while i < len(input):
    var buf = ""

    if isAlphaAscii(input[i]):
      buf = buf & input[i]
      inc i
      while isAlphaNumeric(input[i]):
        buf = buf & input[i]
        inc i

      if buf == "exit":
        result.add(Token(tokenType: EXIT))
      elif buf == "int":
        result.add(Token(tokenType: INT_TYPE))
      elif buf == "void":
        result.add(Token(tokenType: VOID_TYPE))
      else:
        result.add(Token(tokenType: IDENT, value: some(buf)))

      buf = ""
    elif isDigit(input[i]):
      buf = buf & input[i]
      inc i
      while isDigit(input[i]):
        buf = buf & input[i]
        inc i

      result.add(Token(tokenType: INT_LITERAL, value: some(buf)))
      buf = ""
    elif input[i] == '(':
      result.add(Token(tokenType: OPEN_PAREN))
      inc i
    elif input[i] == ')':
      result.add(Token(tokenType: CLOSE_PAREN))
      inc i
    elif input[i] == ';':
      result.add(Token(tokenType: SEMICOLON))
      inc i
    elif input[i] == '{':
      result.add(Token(tokenType: OPEN_CURLY))
      inc i
    elif input[i] == '}':
      result.add(Token(tokenType: CLOSE_CURLY))
      inc i
    elif isSpaceAscii(input[i]):
      inc i
    else:
      raise newException(Exception, "Tokenization failed!")