Docker+vue+go+mysqlにてデプロイの練習

目的
1.Herokuでデプロイ
2.AWSでデプロイ

工程
1.Docker+go+mysqlにてデータ登録できるかの確認
*注意*
1.gormのPROTOCOL設定は以下のようにすること
PROTOCOL := "tcp(mydqlのコンテナ名:mysqlコンテナポート番号)"

2.goのコンテナのポート番号とginのポート番号は統一することでうまく動いた。
今回の場合は以下の通り
docker: 3000:3000
main.go: r.Run(":3000")

2.Docker+go+mysql+vueにて簡単なTodoアプリを作成

3.Herokuデプロイ

4.AWSデプロイ

今後の使用方法
環境構築のベースにすること
