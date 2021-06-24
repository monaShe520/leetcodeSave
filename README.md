# leetcodeSave

- [1.两数之和#1](https://leetcode-cn.com/problems/two-sum/)
    - 题目
        - 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

        - 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

        - 你可以按任意顺序返回答案。
    - 代码
        - [我的解答](https://github.com/monaShe520/leetcodeSave/blob/9daedbc57d19051041129fe9cad383f2bbb09f52/code/leetcode_1_two-sum/main.go#L8)

- [2.两数相加#2](https://leetcode-cn.com/problems/add-two-numbers/)
    - 题目
        - 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
        - 请你将两个数相加，并以相同形式返回一个表示和的链表。
        - 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
    - 代码
        - [我的解答](https://github.com/monaShe520/leetcodeSave/blob/fe56802f9fc2e5ec231499761e779e0bcedd0765/code/leetcode_2_add-two-numbers/main.go#L21)
- [3.无重复字符的最长子串#3](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)
    - 题目
        - 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
    - 代码
        - 我的解答
            - [二层循环-简单](https://github.com/monaShe520/leetcodeSave/blob/36b50675d3a3ea567e0884279bbf2d4e80c01654/code/leetcode_3_longest-substring-without-repeating-characters/main.go#L27)
            - [更优解-todo]
- [4.寻找两个正序数组的中位数#4](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)
    - 题目
        - 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
    - 代码
        - 我的解答
            - [m+n](https://github.com/monaShe520/leetcodeSave/blob/3be596e1fb64b60543ba476e3320613aaeeb6250/code/leetcode_4_median-of-two-sorted-arrays/main.go#L32)
- [5.最长回文子串#5](https://leetcode-cn.com/problems/longest-palindromic-substring/)
    - 题目
        - 给你一个字符串 s，找到 s 中最长的回文子串。
    - 代码
        - 我的解答
            - [n^2中心扩散法](https://github.com/monaShe520/leetcodeSave/blob/aabcdd7dc6d26b81563bf19a13787b7c3524f18a/code/leetcode_5_longest-palindromic-substring/main.go#L22)
- [11.盛最多水的容器#11](https://leetcode-cn.com/problems/container-with-most-water/)
    - 题目
        - 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
    - 代码
        - 我的解答
            - n^2冒泡超时
            - [n双指针法](https://github.com/monaShe520/leetcodeSave/blob/f6365d5a5c212a34fbd57d677fca53a493872050/code/leetcode_11-container-with-most-water/main.go#L11)

- [15.三数之和#15](https://leetcode-cn.com/problems/3sum/)
    - 题目
        - 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
        - 注意：答案中不可以包含重复的三元组。
    - 代码
        - [n^2双指针法](https://github.com/monaShe520/leetcodeSave/blob/8d826a551660d017fe8bd4343bbf7549cb490113/code/leetcode_15_3sum/main.go#L7)

- [17.电话号码的字母组合#17](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
    - 题目
        - 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
        - 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
    - 代码
        - [递归](https://github.com/monaShe520/leetcodeSave/blob/1daf7d4ae30d99e50589f783936a0ff8518eec6e/code/leetcode_17_letter-combinations-of-a-phone-number/main.go#L21)

- [19.删除链表的倒数第N个节点#19](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)
    - 题目
        - 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
        - 进阶：你能尝试使用一趟扫描实现吗？
    - 代码
        - [n 计算链表长度](https://github.com/monaShe520/leetcodeSave/blob/96b98b734a873561167dae544bdd0e0eeef1debc/code/leetcode_19_remove-nth-node-from-end-of-list/main.go#L17)
        - [n 快慢指针](https://github.com/monaShe520/leetcodeSave/blob/96b98b734a873561167dae544bdd0e0eeef1debc/code/leetcode_19_remove-nth-node-from-end-of-list/main.go#L42)

- [20.有效的括号#20](https://leetcode-cn.com/problems/valid-parentheses/)
    - 题目
        - 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
        - 有效字符串需满足：
        - 左括号必须用相同类型的右括号闭合。
        - 左括号必须以正确的顺序闭合。
    - 代码
        - 我的解答
            - [n](https://github.com/monaShe520/leetcodeSave/blob/b997a1ab12e8c25aacffbe596fac748dcbb42557/code/leetcode_20_valid-parentheses/main.go#L54)
- [21.合并两个有序链表#21](https://leetcode-cn.com/problems/merge-two-sorted-lists/)
    - 题目
        - 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
    - 代码
        - 我的解答
            - [m+n](https://github.com/monaShe520/leetcodeSave/blob/f73b9f33ab3fcea87d7524d260aa129224ff999f/code/leetcode_21_merge-two-sorted-lists/main.go#L13)