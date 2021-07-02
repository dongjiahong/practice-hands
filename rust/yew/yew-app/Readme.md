## 安装依赖
```sh
cargo install trunk wasm-bindgen-cli
```

## 修改Cargo.toml文件
```toml
[package]
name = "frontend-yew"
version = "0.1.0"
authors = ["我是谁？"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
wasm-bindgen = "0.2.74"
yew = "0.18.0"
```

## 编译
```sh
trunk build
```
会生成dist目录，里面有wasm文件即，rust代码编译为WebAssemble格式的成果
而dist/index.html也是对index.html文件的编译成果

## 运行
```sh
trunk serve
```
