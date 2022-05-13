#![forbid(unsafe_code)]
#![deny(private_in_public)]
#![warn(
    clippy::all,
    clippy::dbg_macro,
    clippy::todo,
    clippy::empty_enum,
    clippy::enum_glob_use,
    clippy::unused_self,
    clippy::needless_continue,
    clippy::needless_borrow,
    clippy::match_wildcard_for_single_variants,
    clippy::if_let_mutex,
    clippy::mismatched_target_os,
    clippy::match_on_vec_items,
    clippy::imprecise_flops,
    clippy::suboptimal_flops,
    clippy::lossy_float_literal,
    clippy::fn_params_excessive_bools,
    clippy::inefficient_to_string,
    clippy::macro_use_imports,
    clippy::option_option,
    clippy::unnested_or_patterns,
    clippy::str_to_string,
    clippy::cast_lossless,
    clippy::implicit_clone,
    clippy::unused_async,
    clippy::redundant_closure_for_method_calls,
    rust_2018_idioms,
    future_incompatible,
    nonstandard_style,
    missing_debug_implementations
)]

use clap::Parser;
use color_eyre::Result;
use config::{Config, ConfigKey};
use once_cell::sync::Lazy;
use tracing::error;
use tracing_error::ErrorLayer;
use tracing_subscriber::prelude::*;
use tracing_subscriber::{fmt, EnvFilter};

mod api;
mod cmd;
mod config;

pub static VERSION: Lazy<String> = Lazy::new(|| {
    let mut version = format!("v{}", env!("CARGO_PKG_VERSION"));
    if let Some(hash) = option_env!("GIT_SHORT_HASH") {
        version += &format!(" ({})", hash);
    }

    version
});

#[derive(Debug, Parser)]
#[clap(version = &VERSION[..], about, rename_all = "snake_case")]
struct Args {
    /// Verbosity level
    #[clap(short, long, parse(from_occurrences))]
    verbose: u8,

    #[clap(subcommand)]
    subcommand: Subcommand,
}

#[derive(Debug, Parser)]
enum Subcommand {
    /// Finds systems close to others, optionally with constraints
    #[clap(trailing_var_arg = true)]
    CloseTo {
        #[clap(multiple_values = true)]
        args: Vec<String>,
    },

    /// Get / set config keys
    Config {
        #[clap(subcommand)]
        subcommand: ConfigSubcommand,
    },

    /// Returns the coordinates of given systems
    Coords,

    /// Finds the distance between two or more systems
    Distance,

    /// Finds the optimal order to visit a set of stations, and can produce full routes between systems
    Edts,

    /// Searches for systems and stations by name, including wildcards
    Find,

    /// Determines the amount of fuel used by a series of jumps
    FuelUsage,

    /// Gives an estimate of good plot distances in the galactic core
    Galmath,
}

#[derive(Debug, Parser)]
enum ConfigSubcommand {
    /// Get config key
    Get { key: ConfigKey },

    /// Set config key to value
    Set { key: ConfigKey, value: String },
}

fn main() -> Result<()> {
    color_eyre::install()?;
    let args = Args::parse();

    let verbose = args.verbose;
    let filter = match verbose {
        #[cfg(debug_assertions)]
        0 | 1 | 2 => format!("{}=debug", env!("CARGO_PKG_NAME")),

        #[cfg(not(debug_assertions))]
        0 => format!("{}=info", env!("CARGO_PKG_NAME")),
        #[cfg(not(debug_assertions))]
        1 | 2 => format!("{}=debug", env!("CARGO_PKG_NAME")),

        3 => format!("{}=trace", env!("CARGO_PKG_NAME")),
        _ => "trace".into(),
    };

    let filter = EnvFilter::new(filter);
    let fmt = fmt::layer().with_target(verbose >= 2);
    tracing_subscriber::registry()
        .with(filter)
        .with(fmt)
        .with(ErrorLayer::default())
        .init();

    let home_dir = match dirs::home_dir() {
        Some(dir) => dir,

        None => {
            error!("Failed to resolve your home directory!");
            std::process::exit(1);
        }
    };

    let config_path = home_dir.join(".config").join("redts.toml");
    let mut config = Config::load_path(&config_path)?;
    config.url = Some("abc".into());

    match args.subcommand {
        Subcommand::Config { subcommand } => match subcommand {
            ConfigSubcommand::Get { key } => cmd::config_get(&config, key)?,
            ConfigSubcommand::Set { key, value } => cmd::config_set(&mut config, key, value)?,
        },

        Subcommand::CloseTo { args } => todo!(),
        Subcommand::Coords => todo!(),
        Subcommand::Distance => todo!(),
        Subcommand::Edts => todo!(),
        Subcommand::Find => todo!(),
        Subcommand::FuelUsage => todo!(),
        Subcommand::Galmath => todo!(),
    }

    config.save()?;
    Ok(())
}
