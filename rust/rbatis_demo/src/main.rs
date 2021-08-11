
#[macro_use]
extern crate lazy_static;

use fast_log::init_log;
use log::info;
use rbatis::crud::CRUD;
use rbatis::crud_table;
use rbatis::rbatis::Rbatis;

lazy_static! {
    static ref RB: Rbatis = Rbatis::new();
}

#[crud_table]
#[derive(Clone, Debug)]
pub struct TSetting {
    pub id: Option<i32>,
    pub s_key: Option<String>,
    pub s_value: Option<String>,
    pub cfg_tip: Option<String>,
    pub created: Option<i32>,
    pub updated: Option<i32>,
    pub deleted: Option<i32>,
}

#[crud_table(table_name:t_notice)]
#[derive(Clone, Debug)]
pub struct TNotice {
    pub id: Option<i32>,
    pub title: Option<String>,
    pub summary: Option<String>,
    pub cover: Option<String>,
    pub content: Option<String>,
    pub created: Option<i32>,
    pub updated: Option<i32>,
    pub deleted: Option<i32>,
}

#[tokio::main]
async fn main() {
    init_log("requests.log", 1000, log::Level::Info, None, true).unwrap();
    let uri = "mysql://root:root123@172.18.3.2:3306/db_official";
    info!("{}", uri);

    RB.link(uri).await.unwrap();
    info!("link db success");

    // 单查一个
    let w = RB.new_wrapper().eq("id", 1);
    let r: Option<TSetting> = RB.fetch_by_wrapper(&w).await.unwrap();
    println!("{:#?}", r);
    println!("====================================");
    // 查所有
    let r: Vec<TSetting> = RB.fetch_list().await.unwrap();
    println!("{:?}", r);
}


