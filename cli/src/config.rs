use std::fs;
use std::path::{Path, PathBuf};

use color_eyre::Result;
use paste::paste;
use serde::{Deserialize, Serialize};
use strum_macros::EnumString;

macro_rules! define_config {
    (pub struct $name:ident { $($fname:ident : $ftype:ty),* }) => {
        paste! {
            #[derive(Debug, Default, Deserialize, Serialize)]
            pub struct $name {
                #[serde(skip)]
                file_path: PathBuf,

                $($fname : $ftype),*
            }

            #[derive(Debug, EnumString)]
            #[strum(serialize_all = "snake_case")]
            pub enum [<$name Key>] {
                $([<$fname:camel>]),*
            }
        }
    }
}

define_config! {
    pub struct Config {
        url: Option<String>,
        username: Option<String>,
        token: Option<String>
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
