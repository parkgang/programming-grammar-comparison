async function task(sec, content) {
  await new Promise((r) => {
    setTimeout(() => {
      console.log("\t" + content);
      r();
    }, sec);
  });
}

console.log("비동기 작업 시작");

Promise.all([
  task(200, "1번째 비동기 작업"),
  task(1000, "2번째 비동기 작업"),
  task(10, "3번째 비동기 작업"),
  task(500, "4번째 비동기 작업"),
  task(600, "5번째 비동기 작업"),
]).then(() => {
  console.log("비동기 작업 완료! 프로그램을 종료합니다.");
});
