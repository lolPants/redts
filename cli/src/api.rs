use color_eyre::Result;
use once_cell::sync::Lazy;
use reqwest::blocking::{Client, ClientBuilder};
use reqwest::StatusCode;
use tracing::{error, info};

use crate::config::Config;
use crate::VERSION;

const NAME: &str = env!("CARGO_PKG_NAME");
static CLIENT: Lazy<Client> = Lazy::new(|| {
    ClientBuilder::new()
        .user_agent(format!("{}/{}", NAME, *VERSION))
        .build()
        .unwrap()
});

pub fn call_api(config: &Config, script: &'static str, args: Vec<String>) -> Result<()> {
    if config.username.is_some() && config.token.is_none() {}

    let url = match &config.url {
        Some(url) => url.trim_end_matches('/'),
        None => {
            error!("url is unset");
            info!("use: redts config set url <value>");

            std::process::exit(1);
        }
    };

    let url = format!("{url}/api/{script}");
    let mut request = CLIENT
        .get(url)
        .header("Content-Type", "text/plain")
        .body(args.join(" "));

    if let Some(username) = config.username.as_deref() {
        if config.token.is_none() {
            error!("missing auth token");
            std::process::exit(1);
        }

        let token = config.token.as_deref().unwrap();
        let bearer = format!("{username}:{token}");

        request = request.bearer_auth(bearer);
    }

    let resp = request.send()?;
    if resp.status() == StatusCode::UNAUTHORIZED {
        error!("invalid credentials");
        std::process::exit(1);
    }

    let body = resp.error_for_status()?.text()?;
    println!("{body}");

    Ok(())
}
