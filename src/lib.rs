pub mod app;
pub mod routes;
// pub mod middlewares;
pub mod components {
    pub mod navbar;
    pub mod alert;
}
pub mod pages {
    pub mod index;
    pub mod login;
    pub mod signup;
}

pub use app::App;
pub use routes::Routes;
pub use components::{navbar::Navbar, alert::{Warning, Error, Success, Info}};
// Pages 
pub use pages::{index::IndexPage, login::LoginPage, signup::SignupPage};

// middlewares
// pub use middlewares::is_required;