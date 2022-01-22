extern crate proc_macro;
use proc_macro::TokenStream;

use quote::quote;
use syn;

// 函数宏
#[proc_macro]
pub fn make_hello(item: TokenStream) -> TokenStream {
    let name = item.to_string();
    let hell = "Hello ".to_string() + name.as_ref();
    let fn_name =
        "fn hello_".to_string() + name.as_ref() + "(){ println!(\"" + hell.as_ref() + "\"); }";
    fn_name.parse().unwrap()
}

// 属性宏(两个参数)
#[proc_macro_attribute]
pub fn log_att(attr: TokenStream, item: TokenStream) -> TokenStream {
    println!("Attr: {}", attr.to_string());
    println!("Item: {}", item.to_string());
    item
}

// 派生宏
#[proc_macro_derive(Hello)]
pub fn hello_derive(input: TokenStream) -> TokenStream {
    let ast: syn::DeriveInput = syn::parse(input).unwrap();
    let name = ast.ident;
    let gen = quote! {
        impl #name {
            fn hello_macro_derive() {
                println!("Hello proc macro");
            }
        }
    };
    gen.into()
}

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}
