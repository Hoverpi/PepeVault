use sycamore::prelude::*;
use sycamore_router::{Router, HistoryIntegration, Capture, Integration};
use crate::Routes;
use crate::pages::{index::IndexPage, login::LoginPage, signup::SignupPage};

#[component]
pub fn App() -> View {
    let is_authenticated = create_signal(false);

    view! {
        Router(
            integration=HistoryIntegration::new(),
            view=|route: ReadSignal<Routes>| {
                view! {
                    div () {
                        (match route.get_clone() {
                            Routes::Index => view! { IndexPage() },
                            Routes::Login => view! { LoginPage() },
                            Routes::Signup => view! { SignupPage() },
                            // Routes::Vault => require_auth(&is_authenticated || view! { VaultPage() }),
                            Routes::NotFound => view! { "404 Not Found" }
                        })
                    }
                }
            }
        )        
    }
}