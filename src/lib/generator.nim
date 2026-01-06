import parser
import std/strformat

proc generateIntLiteral*(node: IntLiteralNode): string =
    result = fmt"{node.value}"

proc generateExit*(node: ExitNode): string =
    result = fmt"main:{'\n'}" & fmt"    li $v0, {generateIntLiteral(node.exit_code)}{'\n'}" & fmt"    jr $ra{'\n'}"