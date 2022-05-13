use color_eyre::Result;
use tracing::info;

use crate::config::{Config, ConfigKey};

pub fn cmd_config_set(config: &mut Config, key: ConfigKey, value: String) -> Result<()> {
    info!("{key} = \"{value}\"");
    config.set_value(&key, Some(value));

    Ok(())
}
