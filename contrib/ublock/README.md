# Android 用 Stack Overflow の機械翻訳サイトの除外用フィルタ

uBlacklist が対応している5つの検索エンジン (Google, Bing, DuckDuckGo, Ecosia, Startpage) の検索結果から Stack Overflow の機械翻訳サイトを除外します。Firefox for Android にインストールした uBlock Origin で購読することを想定しています。

このリポジトリは、[arosh/ublacklist-stackoverflow-translation](https://github.com/arosh/ublacklist-stackoverflow-translation) のフォークです。

[購読する](https://subscribe.adblockplus.org?location=https%3A%2F%2Fraw.githubusercontent.com%2Fhirorpt%2Fubo-stackoverflow-translation%2Fmaster%2Fandroid%2FuBlockOrigin.txt&title=uBlacklist%20Stack%20Overflow%20Translation) [中身を見る](https://raw.githubusercontent.com/hirorpt/ubo-stackoverflow-translation/master/android/uBlockOrigin.txt)

(このリポジトリは試験的な目的で公開しています。フィルタ名の変更やリポジトリの削除は予告なく行われるため、機能テストの範疇を超えて実運用する場合はフォークしてください。)

## Known limitations

- それぞれモバイル版の、Google, Bing, DuckDuckGo, Ecosia, Startpage の通常検索と、 Google, Bing, DuckDuckGo の画像検索をサポートしています。デスクトップサイトや機械翻訳サイトがまず出現しないであろう特殊検索 (ニュース検索など) については対応していません。ただし巻き添えで消えているものもあります。
- URLの部分一致で判定しているため、例えば機械翻訳サイトに対するはてなブックマークや Similarweb なども除外されます(DuckDuckGoの通常検索とBingの画像検索除く)。また`example.com`を除外するフィルタは`hoge-example.com`も除外するでしょう。
    - 消えては不都合なサイトがあれば個別対応します。汎用的な対策はパフォーマンスが下がる恐れがあるためやりません。
- uBlock Origin、AdGuard 拡張機能、HTTPSフィルタリングを有効にした AdGuard for Android で購読できます。他の広告ブロックアプリやブラウザネイティブのブロック機能とは互換性がありません。

## Tips

- uBlock Origin ではリクエストログから</>ボタンをタップしてDOMインスペクタを表示することで、何が消えているか視覚的に確認することができます。また Google と Bing ではデスクトップサイトを表示することでも簡易的に調べることができます。
