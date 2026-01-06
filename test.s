another_label:
    li $v0, 1
    li $a0, 15
    syscall
    jr $ra

main:
    push $ra
    jal another_label
    pop $ra

    li $v0, 0
    jr $ra
