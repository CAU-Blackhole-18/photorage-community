### 보완해야할 내용
* Error Handling
* 인터페이스 등에 대한 보완
* Golang에서 좋은 로깅 시스템 확인
    - https://medium.com/@123notactive/golang-logging-%EC%A0%84%EB%9E%B5-f6bdbb7bad14
* Kafka Error Handling 보강 필요
    - 장애 상황 등에 대한 보강
    - https://www.popit.kr/golang%EC%97%90%EC%84%9C-%EC%B9%B4%ED%94%84%EC%B9%B4-%EC%BB%A8%EC%8A%88%EB%A8%B8-%EA%B7%B8%EB%A3%B9%EA%B3%BC-%EC%9E%AC%EC%8B%9C%EB%8F%84%EB%A1%9C-%EA%B2%B0%EA%B3%BC%EC%A0%81-%EC%9D%BC%EA%B4%80/
* Repository 레벨부터의 Error 처리 추가
* Repository -> Service -> Controller 까지의 Error Propagation을 어떻게 해야하는가?
    - Spring 에서의 ExceptionHandler를 통한 HTTP Status code 리턴과 같은 방안이 없을까?
    - 그러면, Repository error -> Service 전파 -> Controller에서 catch해서 HTTP Status Code를 내려주어야 하는가?