# 『プログラミング言語Go』正誤表
http://www001.upp.so-net.ne.jp/yshibata/#ERRATA

# 『プログラミング言語Go』のソースコード
https://github.com/adonovan/gopl.io/

# Go言語 - os.Argsでコマンドラインパラメータを受け取る
- os.Args[1:]の意味がわからなくて検索

http://blog.y-yuki.net/entry/2017/04/30/000000



# Goの配列について
- インデックスがわからなくて検索
http://cuto.unirita.co.jp/gostudy/post/go_array/

# ストリームとは？
http://www.geocities.jp/m_hiroi/golang/abcgo13.html
データの流れを抽象化したデータ構造を「ストリーム (stream) 」と呼びます。
たとえば、ファイル入出力はストリームと考えることができます。
また、配列 (スライス) を使ってストリームを表すこともできます。
ただし、単純な配列では有限個のデータの流れしか表すことができません。
ところが、「遅延評価」を用いると擬似的に無限個のデータを表すことができます。これを「遅延ストリーム」と呼びます。

# ioutil.ReadFile
http://cuto.unirita.co.jp/gostudy/post/standard-library-file-io/
ioutil.ReadFile関数はファイルからすべてのデータを入力する関数

# RGBA
goでの使い方
http://twinbird-htn.hatenablog.com/entry/2017/04/23/001743
カラーチャート
https://www.scollabo.com/banban/lectur/websafe.html

# エラー
too few values in struct initializer
https://qiita.com/cotrpepe/items/b8e7f70f27813a846431

# ioutil.ReadAll
https://waman.hatenablog.com/entry/2017/10/05/132727
ioutil.ReadAll 関数は、引数の io.Reader から内容を全て読み込んでバイトスライスとして返します：

# 排他制御にsync.Mutex
https://qiita.com/t-mochizuki/items/80dc959b4221c53f3c40

# flagパッケージ
使い方
https://qiita.com/TsuyoshiUshio@github/items/76d53460e9404c3ace33
Go言語でサクッとコマンドラインツールをつくる
https://medium.com/eureka-engineering/go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%82%B5%E3%82%AF%E3%83%83%E3%81%A8%E3%82%B3%E3%83%9E%E3%83%B3%E3%83%89%E3%83%A9%E3%82%A4%E3%83%B3%E3%83%84%E3%83%BC%E3%83%AB%E3%82%92%E3%81%A4%E3%81%8F%E3%82%8B-e16e97c00913

# オペランド
意味
https://wa3.i-3-i.info/word13306.html

# レキシカルスコープ
http://ubur1114.hatenablog.com/entry/2017/08/07/204606

# os.EXIT(1)
goの場合明示的にEXITコードをつけてあげる
エラーコードを付けないと、このGoのコマンドを使ってBashとかを書くと
エラーを正しく処理できない

# main関数がreturnするとプログラムが終了する
returnをすると他になんの処理をしていても終了する

# ポインタとアドレス
https://wa3.i-3-i.info/word12814.html

# コンパイラでセミコロンを勝手に入れている
セミコロンはほとんど登場しない
forで出てくるくらい

# Goフォーマット
1 + 1 * 3とかであると
→1 + 1*3になり”*”の方が優先度が高いと判定してフォーマットをかける

# string.Joinは何故はないのか？
予めサイズを計算した上で、値を入れていくため

# stringsにエラーを返す時はnilは入らないのでから文字列""を入れるしか無い、

# 関数名に日本語名をつけると、公開範囲は公開になる。
大文字と小文字の概念が無い。

# for文で使う変数名はiとか短くなるがわかりにくくなる
単一loop分はiで良さそう。
ネストした場合はiに意味があるはず。その時は意味のわかる変数名をつける。

# ヒープ・スタック
『ヒープ領域』
アプリケーションやOSで動的に割り当てたり解放するもの
プログラムで一時的に必要になるメモリで、例えばファイルを読み出すときに読みだしたファイル内容を置いておいたり、ネットワークでデータを送受信する時にデータを置いておく時に使う
『スタック領域』
コンパイラやOSが割り当て、アプリケーションでは自由に操作できない領域
プログラムが内部的にデータを保存しておく必要がある場合に、スタック領域が使用される
https://www.uquest.co.jp/embedded/learning/lecture16.html