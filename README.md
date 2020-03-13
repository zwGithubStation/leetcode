# leetcode
source code of leetcode problems

**1 两数之和**  
https://leetcode-cn.com/problems/two-sum/    
注意点：  
1.C的base哈希实现及冲突处理  
2.负数的取余动作存在堆内存溢出的可能，使用取绝对值方式避免  
3.使用-fsanitize=address编译参数检测堆内存是否有溢出   
4.使用<limits.h>的INT_MAX/INT_MIN做初值，可以不用另设初始状态值   
5.-fno-omit-frame-pointer编译选项优化的基本原理   

**2 两数相加**  

**3 无重复字符的最长子串**

**5 最长回文子串**  

**6 Z字形变换**   

**7 整数反转**  
https://leetcode-cn.com/problems/reverse-integer/   
注意点：    
1.边界判断  
2.传入函数的&param，注意未付初值时函数内是否不做处理，返回无效值  
3.分析简单策略　   

**8 字符串转换整数**   

**9 回文数**  
https://leetcode-cn.com/problems/palindrome-number/   
注意点：   
1.简洁方式实现   
2.有效剪纸避免错误覆盖(e.g.100000) 

**11 盛最多水的容器**  

**14 最长公共前缀**  
https://leetcode-cn.com/problems/longest-common-prefix/   
注意点：   
1.初值设定与否，会影响特殊情况下的赋值逻辑   
**2.官网的诸多理论方法**  

**15 三数之和**  

**16 最接近的三数之和**  

**17 电话号码的字母组合**   

**18 四数之和**  

**19 删除链表的倒数第N个节点**   

**20 有效括号字符串**  
https://leetcode-cn.com/problems/valid-parentheses/   
注意点：   
1.stack简单应用的清晰实现代码(逻辑要清晰)   

**21 合并有序lists**  
https://leetcode-cn.com/problems/merge-two-sorted-lists/   
注意点：  
1.算法实现要与构思一致，条例清晰，边界考虑完整    
**2.官网的诸多理论方法**  

**22 括号生成**  

**24 两两交换链表中的节点** 

**26 删除排序数组中的重复项**  
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/  
注意点：  
1.还是边界条件的满足    
2.先实例演算清楚，再着手实现  

**28 实现strStr()**   
https://leetcode-cn.com/problems/implement-strstr/   
注意点：  
**1.C语言的标准库函数(e.g.strncmp的实现及其性能点)**   

**29 两数相除**	

**53 最大子序和**  
https://leetcode-cn.com/problems/maximum-subarray/   
注意点：  
1.三种方式的算式理解   

**58 最后一个单词的长度**   
https://leetcode-cn.com/problems/length-of-last-word/   
无特殊注意点，再次强调变量的初值设定问题   

**66 加一**   
https://leetcode-cn.com/problems/plus-one/  
无特殊注意点   

**69 x的平方根**  
https://leetcode-cn.com/problems/sqrtx/   
注意点：  
**1.四种方法(baseline,baseline-binary,recursive,Newton)细品其中思路**  
2.注意边界条件跳出的判断(e.g. recursive中最终返回那个iter)   

**70 爬楼梯**  


**83 删除排序链表中的重复元素**

**88 合并两个有序数组**

**100 相同的树**

**101对称二叉树**  



