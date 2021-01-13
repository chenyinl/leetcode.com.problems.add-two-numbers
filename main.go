/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l *ListNode; // list location
    var p *ListNode; // list head
    stop1 :=0;
    stop2 :=0;
    v1 := l1.Val;
    v2 := l2.Val;
    nextadd := 0;
    thisadd := 0;
    var v int;
    for (stop1==0 || stop2==0 || nextadd ==1){
        thisadd = nextadd;
        nextadd = 0;
        v = v1+v2+thisadd;
        if(v>9){
            v=v-10;
            nextadd = 1;
        }
        l = addToList(l, v);
        if(p==nil){
            p=l; // the head of list
        }
        if(stop1==1){
        }else if(l1.Next == nil){
            stop1 = 1;
            v1=0;
        }else{
            l1=l1.Next;
            v1=l1.Val;
        }
        if(stop2 ==1){
        }else if(l2.Next == nil){
            stop2 = 1;
            v2=0;
        }else{
            l2=l2.Next;
            v2=l2.Val;
        }
    }
    return p;
}

func addToList(l *ListNode, v int) *ListNode{
    newNode := ListNode{
        Val: v,
        Next: nil,
    }
    if(l==nil){
        l = &newNode;
    }else{
        l.Next = &newNode;
        l = l.Next;
    }
    return l;
}
