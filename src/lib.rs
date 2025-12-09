pub mod app;
pub mod routes;
// pub mod middlewares;
pub mod components {
    pub mod navbar;
    pub mod art;
}
pub mod pages {
    pub mod index;
    pub mod login;
    pub mod signup;
}

pub use app::App;
pub use routes::Routes;
pub use components::navbar::Navbar;
pub use components::art::RustLockArt;
// Pages 
pub use pages::index::*;

// middlewares
// pub use middlewares::is_required;