(module
    (export "square" (func $square))
    (export "fibWasm" (func $fibWasm))

    (func $square (param i32) (result i32)
        local.get 0
        local.get 0
        i32.mul
    )

    (func $fibWasm (param $n i32) (result i32)
        (if
            (i32.lt_s
                (local.get $n)
                (i32.const 1)
            )
            (return
                (i32.const 0)
            )
        )
        (if
            (i32.lt_s
                (local.get $n)
                (i32.const 2)
            )
            (return
                (i32.const 1)
            )
        )
        (return 
            (i32.add
                (call $fibWasm
                    (i32.sub
                        (local.get $n)
                        (i32.const 2)
                    )
                )
                (call $fibWasm
                    (i32.sub
                        (local.get $n)
                        (i32.const 1)
                    )
                )
            )
        )
    )
)