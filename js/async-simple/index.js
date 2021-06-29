async function WorkDo(sec, content) {
  await new Promise((r) => {
    setTimeout(() => {
      console.log("\t" + content);
      r();
    }, sec);
  });
}

console.log("js 비동기 작업 시작");

Promise.all([
  WorkDo(200, "1번째 비동기 작업"),
  WorkDo(1000, "2번째 비동기 작업"),
  WorkDo(10, "3번째 비동기 작업"),
  WorkDo(500, "4번째 비동기 작업"),
  WorkDo(600, "5번째 비동기 작업"),
]).then(() => {
  console.log("비동기 작업 완료! 프로그램을 종료합니다.");
});
