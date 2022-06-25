package main

class A {
    public static Queue‹String> levelSerial (Node head)
    Queue<String>
            ans
new LinkedList‹>();
if (head == null) {
        ans.add (null) ;
    } else {
        ans.add (String. value0f(head. value)) ;
        Queue<Node> queue
        new LinkedList<Node>():
        queue.add (head);
        while (!queue.isEmpty()) {
            head
                    = queue.pb110;
            if (head. left != null) {
                ans.add (String. value0f(head.left.value) );
                queue.add(head.left);
            }
            else f
            ans.add (null) ;
        }
        if (head.right != null) {
            ans.add(String.valueOf(head.right.value));
            queue.add(head.right);
        } else {
            ans.add(null);
        }
    }
return ans;
}
}