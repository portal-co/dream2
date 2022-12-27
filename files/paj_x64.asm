start:
mov %eax, 57
syscall
cmp %eax, 0
je done

bin:
.incbin "exe"

done:
ret