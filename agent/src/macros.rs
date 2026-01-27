#[macro_export]
macro_rules! info {
    () => {
        println!("{}", Local::now().format("%Y/%m/%d %H:%M:%S"));
    };
    ($($arg:tt)*) => {
        println!(
            "{} {}",
            Local::now().format("%Y/%m/%d %H:%M:%S"),
            format_args!($($arg)*)
        );
    };
}