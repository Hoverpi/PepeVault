use sycamore::prelude::*;

#[component]
pub fn Warning(msg: String) -> View {
    // sanitize input
    let sanitize: String = msg.chars()
                              .filter(|c| {
                                  c.is_ascii_alphabetic() || *c == ' '
                              })
                              .collect();
    let warning_msg = format!("Warning: {}!", &sanitize);
    
    view! {
        div (class="alert alert-warning") {
            img(src="/public/icons/triangle-exclamation.svg", alt="", style="color: #f5c211;") {}
            span() { (warning_msg) }
        }
    }
}

#[component]
pub fn Error(msg: String) -> View {
    // sanitize input
    let sanitize: String = msg.chars()
                              .filter(|c| {
                                  c.is_ascii_alphabetic() || *c == ' '
                              })
                              .collect();
    let error_msg = format!("Error! {}.", &sanitize);
    
    view! {
        div (class="alert alert-warning") {
            img(src="/public/icons/circle-exclamation.svg", alt="", style="color: #c01c28;") {}
            span() { (error_msg) }
        }
    }
}

#[component]
pub fn Success(msg: String) -> View {
    // sanitize input
    let sanitize: String = msg.chars()
                              .filter(|c| {
                                  c.is_ascii_alphabetic() || *c == ' '
                              })
                              .collect();
    let success_msg = format!("{}!", &sanitize);
    
    view! {
        div (class="alert alert-warning") {
            img(src="/public/icons/circle-check.svg", alt="", style="color: #2ec27e;") {}
            span() { (success_msg) }
        }
    }
}

#[component]
pub fn Info(msg: String) -> View {
    // sanitize input
    let sanitize: String = msg.chars()
                              .filter(|c| {
                                  c.is_ascii_alphabetic() || *c == ' '
                              })
                              .collect();
    let info_msg = format!("{}!", &sanitize);
    
    view! {
        div (class="alert alert-warning") {
            img(src="/public/icons/circle-info.svg", alt="", style="color: #1c71d8;") {}
            span() { (info_msg) }
        }
    }
}