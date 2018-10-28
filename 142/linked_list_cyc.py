# coding: utf-8

class Solution(object):
    def detectCycle1Map(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        step, walked = head, {}
        while step is not None:
            if step not in walked:
                walked[step] = True
            else:
                return True
            step = step.next

        return False

    def detectCycle1TwoPointers(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow, fast = slow.next, fast.next.next
            if fast is slow:
                return True
        return False

    def detectCycle2Map(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        step, walked = head, {}
        while step is not None:
            if step not in walked:
                walked[step] = True
            else:
                return step
            step = step.next

        return None

    def detectCycle2TwoPointers(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow, fast = slow.next, fast.next.next
            if fast is slow:
                meeter = head
                while slow != meeter:
                    slow = slow.next
                    meeter = meeter.next
                return meeter
                
        return None
