use std::fs;
use std::path::{Path, PathBuf};

use color_eyre::Result;
use serde::{Deserialize, Serialize};

#[derive(Debug, Default, Deserialize, Serialize)]
struct Config {
    #[serde(skip)]
    file_path: PathBuf,

    pub url: Option<String>,
    pub username: Option<String>,
    pub token: Option<String>,
}

impl Config {
    #[inline(always)]
    pub fn load(&self) -> Result<Self> {
        Self::load_path(&self.file_path)
    }

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
