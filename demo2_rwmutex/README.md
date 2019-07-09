# RWMutex读写锁Demo

运行结果
```
[root@aidevops demo2_rwmutex]# go run main.go
Read5   -- read enter
Read5   -- read start
Read5   -- reading...0
Read2   -- read enter
Read2   -- read start
Read2   -- reading...0
Write3  -- write enter
Write4  -- write enter
Read1   -- read enter
Read2   -- reading...1
Read5   -- reading...1
Read5   -- reading...2
Read2   -- reading...2
Read2   -- read over
Read5   -- read over
Write3  -- write start
Write3  -- writeing...0
Write3  -- writeing...1
Write3  -- writeing...2
Write3  -- write over
Read1   -- read start
Read1   -- reading...0
Read1   -- reading...1
Read1   -- reading...2
Read1   -- read over
Write4  -- write start
Write4  -- writeing...0
Write4  -- writeing...1
Write4  -- writeing...2
Write4  -- write over
```
可以看到读锁发生的时候, 其他的锁无需等待, 可以并发执行
当写锁发生的时候, 其他的读写都要等待
