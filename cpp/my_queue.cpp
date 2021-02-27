#include<bits/stdc++.h>
using namespace std;

int main(){

    queue<int> q;
    q.push(1);
    q.push(2);
    q.push(3);
    cout<<q.front()<<endl;
    q.pop();
    cout<<q.front()<<endl;
    q.push(4);
    q.push(5);
    cout<<q.empty()<<endl;

    return 0;
}
