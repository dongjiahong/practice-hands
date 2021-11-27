
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
#[structopt( name = "strings",
             author = "xxx",
             about = "strings - Let's you modify and inspect strings")]
struct CLI {
    #[structopt(long, short, global = true,
                help = "Prints debug information")]
    debug: bool,    // 全局变量
    input: String,  // 全局变量

    #[structopt(subcommand)]
    cmd: SubCommand,
}

#[derive(Debug,StructOpt)]
enum SubCommand {
    #[structopt(name = "mod", about = "Use mod to modify strings")]
    Modify(ModifyOptions),

    #[structopt(name = "insp", about = "Use insp to inspect strings")]
    Inspect(InspectOptions),

    #[structopt(name = "math", about = "Use insp to inspect strings")]
    Maths(Math),
}


#[derive(Debug, StructOpt)]
struct ModifyOptions {
    #[structopt(short, long, help = "Transforms a string to uppercase")]
    upper: bool,
    #[structopt(short, long, help = "Transforms a string to lowercase")]
    lower: bool,
    #[structopt(short, long, help = "Reverses a string")]
    reverse: bool,
    #[structopt(short="pref", long, help = "Adds a prefix to the string", env = "STRINGS__PREFIX")]
    prefix: Option<String>,
    #[structopt(short="suf", long, help = "Adds a suffix to the string", env = "STRINGS__SUFFIX")]
    suffix: Option<String>,
}

#[derive(Debug, StructOpt)]
struct InspectOptions{
    #[structopt(short, long, help = "Count all characters in the string")]
    length: bool,
    #[structopt(short, long, help = "Count only numbers in the given string")]
    numbers: bool,
    #[structopt(short, long, help = "Count all spaces in the string")]
    spaces: bool,
}

#[derive(Debug, StructOpt)]
struct Math {
    #[structopt(subcommand)]
    cmd: Method,
}
#[derive(Debug, StructOpt)]
enum Method {
    #[structopt(name = "add", about = "Use insp to inspect strings")]
    Add(AddOptions),
    #[structopt(name = "mul", about = "Use insp to inspect strings")]
    Mul(MulOptions),
}




#[derive(Debug, StructOpt)]
struct AddOptions {
    #[structopt(short, long, help = "a + b")]
    a: i32,
    #[structopt(short, long, help = "a + b")]
    b: i32,
}

#[derive(Debug, StructOpt)]
struct MulOptions {
    #[structopt(short, long, help = "a + b")]
    a: i32,
    #[structopt(short, long, help = "a * b")]
    b: i32,
}


fn modify(input: &String, debug: bool, args: &ModifyOptions) {
    println!("Modify called for {}", input);
    if debug {
        println!("{:#?}", args);
    }
}

fn inspect(input: &String, debug: bool, args: &InspectOptions) {
    println!("Inspect called for {}", input);
    if debug {
        println!("{:#?}", args);
    }
}

fn add(input: &String, debug: bool, args: &AddOptions) {
    println!("Add called for {}", input);
    if debug {
        println!("{:#?}", args);
    }
    println!{"{} + {} = {}", args.a, args.b, args.a + args.b}
}

fn mul(input: &String, debug: bool, args: &MulOptions) {
    println!("mul called for {}", input);
    if debug {
        println!("{:#?}", args);
    }
    println!{"{} + {} = {}", args.a, args.b, args.a * args.b}
}


fn main() {
    let args = CLI::from_args();
    match args.cmd {
        SubCommand::Inspect(opt) => {
            inspect(&args.input, args.debug, &opt);
        }
        SubCommand::Modify(opt) => {
            modify(&args.input, args.debug, &opt);
        }
        SubCommand::Maths(opt) => {
            match opt.cmd {
                Method::Add(opt) => {
                    add(&args.input, args.debug, &opt)
                }
                Method::Mul(opt) => {
                    mul(&args.input, args.debug, &opt)
                }
            }
        }
    }
}
