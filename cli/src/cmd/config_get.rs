use color_eyre::Result;
use tracing::info;

use crate::config::{Config, ConfigKey};

pub fn cmd_config_get(config: &Config, key: ConfigKey) -> Result<()> {
    let value = config.get_value(&key);
    match value {
        Some(value) => info!("{key} = \"{value}\""),
        None => info!("{key} = unset"),
    }

    Ok(())
}
