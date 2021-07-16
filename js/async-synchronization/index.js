console.log("비동기 동기화 예제 시작");

setTimeout(() => {
  console.log("Async CallBack!");
}, 1);

for (let i = 0; i < 20; i++) {
  console.log("for loop: " + i);
}

console.log("반복문 종료");
