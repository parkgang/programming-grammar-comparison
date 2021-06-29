using System;
using System.Threading.Tasks;

namespace async_simple
{
  class Program
  {
    async static Task Main()
    {
      Console.WriteLine("c# 비동기 작업 시작");

      await Task.WhenAll(new[]
      {
        WorkDo(200, "1번째 비동기 작업"),
        WorkDo(1000, "2번째 비동기 작업"),
        WorkDo(10, "3번째 비동기 작업"),
        WorkDo(500, "4번째 비동기 작업"),
        WorkDo(600, "5번째 비동기 작업"),
      });

      Console.WriteLine("비동기 작업 완료! 프로그램을 종료합니다.");
    }

    public static async Task WorkDo(int sec, string content)
    {
      await Task.Delay(sec);
      Console.WriteLine("\t" + content);
    }
  }
}
