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


**整数反转**  
https://leetcode-cn.com/problems/reverse-integer/   
注意点：    
1.边界判断  
2.传入函数的&param，注意未付初值时函数内是否不做处理，返回无效值  
3.分析简单策略　  

**回文数**  
https://leetcode-cn.com/problems/palindrome-number/   
注意点：   
1.简洁方式实现   
2.有效剪纸避免错误覆盖(e.g.100000)   