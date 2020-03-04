# leetcode
source code of leetcode problems

**两数之和**  
https://leetcode-cn.com/problems/two-sum/    
注意点：  
1.C的base哈希实现及冲突处理  
2.负数的取余动作存在堆内存溢出的可能，使用取绝对值方式避免  
3.使用-fsanitize=address编译参数检测堆内存是否有溢出   
4.使用<limits.h>的INT_MAX/INT_MIN做初值，可以不用另设初始状态值   
5.-fno-omit-frame-pointer编译选项优化的基本原理
