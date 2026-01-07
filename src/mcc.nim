import std/[cmdline, strutils, strformat]
import lib/[tokenizer, parser, generator]

proc main(): void =
  let argv = commandLineParams()

  let f = open(argv[0], FileMode.fmRead)
  let contents = f.readAll()
  defer: f.close()

  try:
    let tokens = tokenize(contents)
    echo tokens
    #let tree = parse(tokens)
    #let generatedContent = generateExit(tree)
    #let f = open(fmt"{argv[0].split('.')[0]}.s", FileMode.fmWrite)
    #f.write(generatedContent)
    #f.close()
  except Exception as e:
    echo e.msg

when isMainModule:
  main()