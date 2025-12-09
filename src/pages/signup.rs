use sycamore::prelude::*;

#[component]
pub fn SignupPage() -> View {
    let username = create_signal(String::new());
    let password = create_signal(String::new());
    let confirm_password = create_signal(String::new());

    let show_password = create_signal(false);
    let show_confirm_password = create_signal(false);
    let on_submit = {
        let username = username.clone();
        let password = password.clone();
        let confirm_password = confirm_password.clone();

        
    }

    view! {
        div(class="min-h-screen flex items-center justify-center bg-base-200 px-4") {
            div(class="w-full max-w-md bg-base-100 p-6 rounded-xl shadow-lg") {
                h1(class="text-2xl font-bold text-center mb-4") { "Sign up" }
                p(class="text-sm text-muted text-center mb-6") {
                    "Welcome — Sign up to continue to PepeVault"
                }

                form(class="space-y-4", on:submit=on_submit) {
                    // Username: input with icon INSIDE the box (left)
                    div(class="form-control") {
                        label(class="label") {
                            span(class="label-text") { "Username" }
                        }

                        // relative wrapper so we can absolutely position the icon
                        div(class="relative") {
                            // Inline user icon — absolute inside the input
                            img(
                                class="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none opacity-80 z-10",
                                src="/public/icons/user.svg",
                                alt="",
                                aria-hidden="true"
                            ) {}
                            input(
                                class="input input-bordered w-full pl-10", // pl-10 leaves room for icon
                                r#type="text",
                                name="username",
                                placeholder="username",
                                required=true,
                                pattern="[A-Za-z][A-Za-z0-9\\-]*",
                                minlength="3",
                                maxlength="30",
                                title="3–30 chars: letters, numbers or dash"
                            ) {}
                        }

                        p(class="text-xs text-muted mt-1") {
                            "3–30 characters — letters, numbers or dash"
                        }
                    }

                    // Password: input with lock icon inside, and an optional show/hide button on the right
                    div(class="form-control") {
                        label(class="label") {
                            span(class="label-text") { "Password" }
                        }

                        div(class="relative") {
                            // Lock icon inside input (left)
                            img (class="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none opacity-80 z-10",
                                src="/public/icons/key.svg",
                                alt="",
                                aria-hidden="true"
                            ) {}

                            input(
                                class="input input-bordered w-full pl-10 pr-10", // left padding for icon, right for toggle button
                                r#type=if show_password.get() { "text" } else { "password" },
                                name="password",
                                placeholder="Password",
                                required=true,
                                minlength="8",
                                pattern="(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).{8,}",
                                title="At least 8 chars, with number, lower & upper case"
                            ) {}

                            // Show/hide password toggle (right). For now it's a visual button; wiring behavior is optional.
                            button(
                                r#type="button",
                                class="absolute right-2 top-1/2 -translate-y-1/2 flex items-center justify-center z-30 opacity-80 hover:opacity-100 group-focus-within:opacity-100",
                                aria-label="Toggle password visibility",
                                on:click=move |_| show_password.set(!show_password.get())
                            ) {
                                (if show_password.get() {
                                    view! {
                                        img (class="w-5 h-5", src="/public/icons/eye.svg", alt="Hide Password") { }
                                    }
                                } else {
                                    view! {
                                        img (class="w-5 h-5", src="/public/icons/eye-slash.svg", alt="Show Password") { }
                                    }
                                })
                            }
                        }

                        p(class="text-xs text-muted mt-1") {
                            "At least 8 characters, including a number, lowercase and uppercase letter"
                        }
                    }

                    // Confirm Password: input with lock icon inside, and an optional show/hide button on the right
                    div(class="form-control") {
                        label(class="label") {
                            span(class="label-text") { "Confirm Password" }
                        }

                        div(class="relative") {
                            // Lock icon inside input (left)
                            img (class="w-5 h-5 absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none opacity-80 z-10",
                                src="/public/icons/key.svg",
                                alt="",
                                aria-hidden="true"
                            ) {}

                            input(
                                class="input input-bordered w-full pl-10 pr-10", // left padding for icon, right for toggle button
                                r#type=if show_confirm_password.get() { "text" } else { "password" },
                                name="confirm_password",
                                placeholder="Confirm Password",
                                required=true,
                                minlength="8",
                                pattern="(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).{8,}",
                                title="At least 8 chars, with number, lower & upper case"
                            ) {}

                            // Show/hide password toggle (right). For now it's a visual button; wiring behavior is optional.
                            button(
                                r#type="button",
                                class="absolute right-2 top-1/2 -translate-y-1/2 flex items-center justify-center z-30 opacity-80 hover:opacity-100 group-focus-within:opacity-100",
                                aria-label="Toggle password visibility",
                                on:click=move |_| show_confirm_password.set(!show_confirm_password.get())
                            ) {
                                (if show_confirm_password.get() {
                                    view! {
                                        img (class="w-5 h-5", src="/public/icons/eye.svg", alt="Hide Password") { }
                                    }
                                } else {
                                    view! {
                                        img (class="w-5 h-5", src="/public/icons/eye-slash.svg", alt="Show Password") { }
                                    }
                                })
                            }
                        }

                        p(class="text-xs text-muted mt-1") {
                            "At least 8 characters, including a number, lowercase and uppercase letter"
                        }
                    }

                    // Buttons
                    div(class="flex flex-col gap-2 mt-2") {
                        a(class="btn btn-ghost w-full", href="/login", role="button") { "Log in" }
                        button(r#type="submit", class="btn btn-primary w-full") { "Sign up" }

                        div(class="flex items-center justify-between mt-2 text-sm") {
                            a(class="link link-hover", href="/forgot") { "Forgot password?" }
                            a(class="link link-hover", href="/help") { "Help" }
                        }
                    }
                } // form
            } // card
        } // page
    }
}