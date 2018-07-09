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