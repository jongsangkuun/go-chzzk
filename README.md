# go-chzzk
비공개 API를 golang으로 라이브러리화

- chat

    - 채팅 규칙, 채팅방 사용 설정 등 채팅방 관련 파싱 함수 정의

- connection

    -  비공식 치지직 API 주소 설정
    -  API의 버전 업그레이드 시 주소 변경 등의 로직 정의

- donate

    - 도네이션 설정 관련 파싱 함수

- live

    - 방송 상태 파싱 함수 정의
    - 스트리머 4인에 대한 테스트 코드 추가
        - live_details
            - 치지직 live_details 정보 요청 로직

        - live_status
            - 치지직 live_status 정보 요청 로직


- models

    - 비공식 API의 데이터 구조 정의
    - 2024/03/12일 발생 데이터를 기준으로 구조체 정의
        - 부정확한 데이터의 경우 interface로 정의

---
Todo

- live chat connection 기능 개발