use sycamore_router::Route;

#[derive(Route, Clone)]
pub enum Routes {
    #[to("/")]
    Index,
    #[to("/login")]
    Login,
    #[to("/signup")]
    Signup,
    // #[to("/vault")]
    // Vault,
    #[not_found]
    NotFound,
}