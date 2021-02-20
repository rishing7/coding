#include<bits/stdc++.h>

using namespace std;

int main(){

    // Declaring vector for integer data type
    vector<int> v_int;

    // Pushing integer into vector from 0 to 4
    for(int i=0; i<5; i++){
        v_int.push_back(i);
    }

    // Iterating vector using for loop
    for(int i=0; i<v_int.size(); i++){
        cout<<v_int[i]<<endl;

    }

    sort(v_int.begin(), v_int.end(), greater<int>());

    // Iterating using vector iterator
    vector<int>::iterator v = v_int.begin();
    while(v != v_int.end()){
        cout<<*v<<endl;
        v++;
    }

    return 0;
}
