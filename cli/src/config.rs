use std::fs;
use std::path::{Path, PathBuf};

use color_eyre::Result;
use paste::paste;
use serde::{Deserialize, Serialize};
use strum_macros::{Display, EnumString};

macro_rules! define_config {
    (pub struct $name:ident { $($fname:ident),* }) => {
        paste! {
            #[derive(Debug, Default, Deserialize, Serialize)]
            pub struct $name {
                #[serde(skip)]
                file_path: PathBuf,

                pub $($fname: Option<String>),*
            }

            #[derive(Debug, Display, EnumString)]
            #[strum(serialize_all = "snake_case")]
            pub enum [<$name Key>] {
                $([<$fname:camel>]),*
            }

            impl $name {
                #[inline(always)]
                pub fn get_value(&self, key: &[<$name Key>]) -> Option<&str> {
                    match key {
                        $(
                            ConfigKey::[<$fname:camel>] => self.$fname.as_deref(),
                        )*
                    }
                }

                #[inline(always)]
                pub fn set_value(&mut self, key: &[<$name Key>], value: Option<String>) {
                    match key {
                        $(
                            ConfigKey::[<$fname:camel>] => self.$fname = value,
                        )*
                    }
                }
            }
        }
    }
}

define_config! {
    pub struct Config {
        url,
        username,
        token
    }
}

impl Config {
    pub fn load_path(path: &Path) -> Result<Self> {
        if !path.exists() {
            let config = Config {
                file_path: path.to_owned(),
                ..Default::default()
            };

            return Ok(config);
        }

        let bytes = fs::read(path)?;
        let mut config: Config = toml::from_slice(&bytes)?;
        config.file_path = path.to_owned();

        Ok(config)
    }

    #[inline(always)]
    pub fn save(&self) -> Result<()> {
        Self::save_path(self, &self.file_path)
    }

    pub fn save_path(data: &Self, path: &Path) -> Result<()> {
        let data = toml::to_string_pretty(data)?;
        fs::write(path, data)?;

        Ok(())
    }
}
