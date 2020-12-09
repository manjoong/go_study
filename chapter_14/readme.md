## 0. 오류의 발단!
1.  prose의 join 패키지는 여러개의 단어들이 매개변수로 들어 올때 **단어 사이에 ,(콤마)** 를 붙여주고 **마지막 단어 사이엔 and**를 붙여주는 함수임
2.  ["my parents", "a rodeo clown", "a prize bull"] 라는 배열이 들어 왔는데, 이럴 경우 **"A photo of my parents, a rodeo clown and a prize bull(부모님이 로데오의 카우보이이자 황소)"** 라는 의미로 잘못 해석됨.
3. 따라서 **마지막 단어 사이에 "and"가 아닌 ", and"** 를 넣어 각 단어를 모두 분리시켜 문제를 해결함
-  --> (A photo of my parents, a rodeo clown, and a prize bull-> 카우보이와 황소와 함께 있는 부모님의 사진) 
4. 하지만 이럴 경우, 배열의 원소가 두개 있때 문장이 이상해짐!
- --> A photo of my parents, and a rodeo clown
5. 단어가 1, 2, 3 개 들어오듯, 다양한 상황을 위한 오류 검출 테스트 코드를 작성 해야한다!! 


## 1. 테스트 파일 작성법
- join_test.go 처럼 _test의 접미사를 가져야 인식한다.
- 테스트 하려는 코드와 반드시 동일한 패키지에 속할 필요는 없으나, 노출되지 않는 타입이나 함수에 접근하려면 동일한 패키지에 속하는 것이 좋다.
- 테스트 함수의 이름은 Test로 시작해야 한다 ex) TestTwoElements
- 테스트 함수는 단일 매개변수로 testing.T 값의 포인터를 받는다. 
--> 테스트의 오류 보고. 원인을 문자열로 받을 수 있음