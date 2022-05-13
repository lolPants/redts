use once_cell::sync::Lazy;
use reqwest::blocking::{Client, ClientBuilder};

use crate::VERSION;

const NAME: &str = env!("CARGO_PKG_NAME");
pub(crate) static CLIENT: Lazy<Client> = Lazy::new(|| {
    ClientBuilder::new()
        .user_agent(format!("{}/{}", NAME, *VERSION))
        .build()
        .unwrap()
});
