use clap::Parser;
use once_cell::sync::Lazy;

static VERSION: Lazy<String> = Lazy::new(|| {
    let mut version = format!("v{}", env!("CARGO_PKG_VERSION"));
    if let Some(hash) = option_env!("GIT_SHORT_HASH") {
        version += &format!(" ({})", hash);
    }

    version
});

#[derive(Debug, Parser)]
#[clap(version = &VERSION[..], about, rename_all = "snake_case")]
enum Subcommand {
    /// Finds systems close to others, optionally with constraints
    CloseTo,

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
    Get,

    /// Set config key value
    Set,
}

fn main() {
    let subcommand = Subcommand::parse();
    dbg!(subcommand);
}
