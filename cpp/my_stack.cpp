#include<bits/stdc++.h>
using namespace std;


int main(){
    stack<int> s;
    s.push(1);
    s.push(2);
    s.push(3);
    cout<<s.top()<<endl;
    s.pop();
    cout<<s.top()<<endl;
    s.push(4);
    s.push(5);
    cout<<"Size of stack is: "<<s.size()<<endl;
    cout<<s.empty()<<endl;


    return 0;
}
