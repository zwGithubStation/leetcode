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
https://leetcode-cn.com/problems/add-two-numbers/  
注意点：  
1.代码稍显冗余，是否有同样逻辑下的简洁实现   
2.是否有其他更优逻辑   
**3.多个指针的赋值相互影响的问题** 

**3 无重复字符的最长子串**  
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/    
注意点：  
**1.滑动窗口和优化滑动窗口的细化理解**  
**2.空间消耗过大 是否直接使用数组**   

**5 最长回文子串**  
https://leetcode-cn.com/problems/longest-palindromic-substring/  
注意点：  
**1.基本方法、动态规划、暴力法的逻辑理解**    
**2.空间由平方到n级的缩减策略理解**  
**3.马拉车算法的理解实现**       

**6 Z字形变换**  
https://leetcode-cn.com/problems/zigzag-conversion/   
注意点： 
1.数学方法的思想理清,边界条件仍要清晰  
2.其他优雅的方法细化理解    

**7 整数反转**  
https://leetcode-cn.com/problems/reverse-integer/   
注意点：    
1.边界判断  
2.传入函数的&param，注意未付初值时函数内是否不做处理，返回无效值  
3.分析简单策略　   

**8 字符串转换整数**  
https://leetcode-cn.com/problems/string-to-integer-atoi/  
注意点：  
1.简单的字符串处理,着重于各个边界条件的判断正确性和异常用例的覆盖性    

**9 回文数**  
https://leetcode-cn.com/problems/palindrome-number/   
注意点：   
1.简洁方式实现   
2.有效剪纸避免错误覆盖(e.g.100000) 

**11 盛最多水的容器**  
https://leetcode-cn.com/problems/container-with-most-water/   
注意点：   
1.双指针法的合理解释理解(见题解之：O(n) 双指针解法：理解正确性、图解原理（C++/Java）)   

**14 最长公共前缀**  
https://leetcode-cn.com/problems/longest-common-prefix/   
注意点：   
1.初值设定与否，会影响特殊情况下的赋值逻辑   
**2.官网的诸多理论方法**  

**15 三数之和**  
https://leetcode-cn.com/problems/3sum/  
注意点：   
**1.排序+寻结果策略的算法理解**  
**2.代码溢出点的排查注意：1.多个去重场景的考虑；2.存储结果内存与数组规模的关系**   
**3.C标准的qsort/bsort的使用和自实现；可重入及线程安全性扩展理解**   

**16 最接近的三数之和**   
https://leetcode-cn.com/problems/3sum-closest/  
注意点：  
和水桶问题相近似  

**17 电话号码的字母组合**  
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/    
注意点：  
**1.快速debug的紧迫性，不要大量使用printf**  
**2.失误点仍常犯：1.进位的时候，地位要变0；2.最后退出时不用走计算函数否则溢出**  
3.其他算法的了解    

**18 四数之和**  

**19 删除链表的倒数第N个节点**   
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/    
注意点：  
简单的链表操作,关注目的是否是常量数组形式记录指针集,关注评论中是否有可借鉴之处      

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
https://leetcode-cn.com/problems/swap-nodes-in-pairs/  
注意点：  
简单的链表操作,关注评论中是否有可借鉴之处  

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
https://leetcode-cn.com/problems/climbing-stairs/   
注意点：  
1.使用公式法时,常规的累积乘方式很容易溢出,根据公式的规律,采取乘一个然后除一个的方式进行计算  
2.(int)A/B 和 (int)(A/B)的概念可不一样呦....    
**3.TO-DO 其他推荐方法细品其中思路**

**83 删除排序链表中的重复元素**  
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/  
注意点：  
简单的测试基本操作；策略明晰则代码简洁，思考为什么提交4次才成功   

**88 合并两个有序数组**  
https://leetcode-cn.com/problems/merge-sorted-array/  
注意点：  
**1.TO-DO 其他推荐方法细品其中思路**

**100 相同的树**  
https://leetcode-cn.com/problems/same-tree/  
注意点：  
**1.基本的树逻辑，关注评论中的可关注点及可优化点**  
**2.递归与迭代**     

**101对称二叉树**  
https://leetcode-cn.com/problems/symmetric-tree/  
注意点：  
**1.基本的树逻辑，关注评论中的可关注点及可优化点**
**2.递归与迭代**     



