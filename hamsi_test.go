package hamsi

import (
    "bytes"
    "testing"
    "encoding/hex"
)

func fromHex(s string) []byte {
    h, _ := hex.DecodeString(s)
    return h
}

type testData struct {
    msg []byte
    md []byte
}

func Test_Hash224_Check(t *testing.T) {
   tests := []testData{
        {
           fromHex(""),
           fromHex("b9f6eb1a9b990373f9d2cb125584333c69a3d41ae291845f05da221f"),
        },
        {
           fromHex("cc"),
           fromHex("8bfa48cf172314d558417877cda9be97825128c531165407fc241040"),
        },
        {
           fromHex("41fb"),
           fromHex("5eabc4770ad6ab30335ca58de088aa234db09258933ba833113a5fa1"),
        },
        {
           fromHex("1F877C"),
           fromHex("15a0b54528fe0f765b50bd340bfb36ae32f106e305aec3b2f42cbec5"),
        },
        {
           fromHex("c1ecfdfc"),
           fromHex("0a3a2bf457fdc3fbeb78dfd423afc35d772ab22bdbe2aeb5af481fa1"),
        },
        {
           fromHex("1f66ab4185ed9b6375"),
           fromHex("ccf83b7f505bfc59fcb40e6eff6dccf54040e30ed914a6fb50af20ee"),
        },
        {
           fromHex("a746273228122f381c3b46e4f1"),
           fromHex("b5ab10136121523143f6e5f94539d9e710a6b7410ac28e14f24aaf0a"),
        },
        {
           fromHex("d8dc8fdefbdce9d44e4cbafe78447bae3b5436102a"),
           fromHex("3f42b2b3154ae54c1da8de3087f4643010f4af632696c61659f44031"),
        },
        {
           fromHex("aecbb02759f7433d6fcb06963c74061cd83b5b3ffa6f13c6"),
           fromHex("6970a83589478a58d9f57dd914b1746ff2269114bbe23a664c03a0a7"),
        },
        {
           fromHex("fac523575a99ec48279a7a459e98ff901918a475034327efb55843"),
           fromHex("b63007f7e3dca1b1af0d7c5711dee2e1aa66680de1faeb74d50942de"),
        },
        {
           fromHex("03a18688b10cc0edf83adf0a84808a9718383c4070c6c4f295098699ac2c"),
           fromHex("ad9b3e9f8d91d4ea0310b861df67b303369b50be40bca7c24c49fb67"),
        },
        {
           fromHex("84fb51b517df6c5accb5d022f8f28da09b10232d42320ffc32dbecc3835b29"),
           fromHex("dbde3421555b22f1469410a2f7fae414484aa6d1562eeb914683a483"),
        },
        {
           fromHex("9f2fcc7c90de090d6b87cd7e9718c1ea6cb21118fc2d5de9f97e5db6ac1e9c10"),
           fromHex("66ffd45e37f72bf74f7143df4fe5d4d99d56109dd86d1374d55f3bef"),
        },
    }

    h := New224()

    for i, test := range tests {
        h.Reset()
        h.Write(test.msg)
        sum := h.Sum(nil)

        if !bytes.Equal(sum, test.md) {
            t.Errorf("[%d] New224 fail, got %x, want %x", i, sum, test.md)
        }

        // =====

        sum2 := Sum224(test.msg)

        if !bytes.Equal(sum2[:], test.md) {
            t.Errorf("[%d] Sum224 fail, got %x, want %x", i, sum2, test.md)
        }
    }
}

func Test_Hash256_Check(t *testing.T) {
   tests := []testData{
        {
           fromHex(""),
           fromHex("750e9ec469f4db626bee7e0c10ddaa1bd01fe194b94efbabebd24764dc2b13e9"),
        },
        {
           fromHex("cc"),
           fromHex("ac2dac2a6ddaf703b7a55745d61b1a16a3d1bf1f74caab265a2e5dbebcf60832"),
        },
        {
           fromHex("41fb"),
           fromHex("2db4f6b7a8e20b28d5d3d536ea23ade6566d4e622e62a108cd52a7a809c469dd"),
        },
        {
           fromHex("1F877C"),
           fromHex("cb596913e691f8654a613e24debf3262e6477fd737d5c422e670e0c75fae7d17"),
        },
        {
           fromHex("c1ecfdfc"),
           fromHex("aeaef9b3b7f8f947fba5fe9bd9886a203110621bc2bca6ac890997aeec69ae0e"),
        },
        {
           fromHex("1f66ab4185ed9b6375"),
           fromHex("271d8b8e833fcac17e0b487ba0f7ee8ddc41a3d34db3390e7ab7e536d71e8564"),
        },
        {
           fromHex("a746273228122f381c3b46e4f1"),
           fromHex("be54b3bda29df786bbd9c460d71c741537bf38cc218357e5fb10b717f8b7f828"),
        },
        {
           fromHex("d8dc8fdefbdce9d44e4cbafe78447bae3b5436102a"),
           fromHex("eed455bc166c62daa2514d5e69c1abc439e8c256c43d0bce222b1ff7336ac1b5"),
        },
        {
           fromHex("aecbb02759f7433d6fcb06963c74061cd83b5b3ffa6f13c6"),
           fromHex("07556a3a185ffe16367e84d8f21fb3a04174a1000d023e12697518b7f942887f"),
        },
        {
           fromHex("7bc84867f6f9e9fdc3e1046cae3a52c77ed485860ee260e30b15"),
           fromHex("8bc5d7e1ab86dc47a2a784ba7823f9ac5906ce79feeb98021c55bfb33226fca4"),
        },
        {
           fromHex("a963c3e895ff5a0be4824400518d81412f875fa50521e26e85eac90c04"),
           fromHex("2be2ef84939028c0c70987d89d58fc927e6142177b13b42f0988005909830468"),
        },
        {
           fromHex("03a18688b10cc0edf83adf0a84808a9718383c4070c6c4f295098699ac2c"),
           fromHex("debde7a74f3533350328f8cd014959dc1a6bf179d7782e5592967f49a867dc74"),
        },
        {
           fromHex("9f2fcc7c90de090d6b87cd7e9718c1ea6cb21118fc2d5de9f97e5db6ac1e9c10"),
           fromHex("6e72391d5be0769c20d92aebee0b1772939e31d521bca1d25f2add261e920ec1"),
        },
    }

    h := New256()

    for i, test := range tests {
        h.Reset()
        h.Write(test.msg)
        sum := h.Sum(nil)

        if !bytes.Equal(sum, test.md) {
            t.Errorf("[%d] New256 fail, got %x, want %x", i, sum, test.md)
        }

        // =====

        sum2 := Sum256(test.msg)

        if !bytes.Equal(sum2[:], test.md) {
            t.Errorf("[%d] Sum256 fail, got %x, want %x", i, sum2, test.md)
        }
    }
}

func Test_Hash384_Check(t *testing.T) {
   tests := []testData{
        {
           fromHex(""),
           fromHex("3943cd34e3b96b197a8bf4bac7aa982d18530dd12f41136b26d7e88759255f21153f4a4bd02e523612b8427f9dd96c8d"),
        },
        {
           fromHex("cc"),
           fromHex("9b299c0b4a6838b5b0f53b0f9c0aea98bbc9c4c9481ec0ec68f344e696f8787de2e08a1404a038c83ac9e121136e8bb8"),
        },
        {
           fromHex("41fb"),
           fromHex("e7394c52238ca2251e51714e790b0ee64a27ebd669cd88f2d564bf17ff704d710ba5f4419dd106a027b16d3decfb3a9a"),
        },
        {
           fromHex("1F877C"),
           fromHex("d8c34c26e4147f706b94923073ee272aef4d024e75cb622288016e38175af79c405cec671f426dc2abef6e4381886e69"),
        },
        {
           fromHex("c1ecfdfc"),
           fromHex("ce995f5ddaa1efca93f01e88cd419a4b7858b6a1624753e3d86998f7a1731dbc1fb9e1461f967d11702f83c5b412c52e"),
        },
        {
           fromHex("1f66ab4185ed9b6375"),
           fromHex("fb38db5b8d9dc1f21a09b379924414260da7fa204cd3de09c95f85fb948e06a8b85f5cfec6dc68ffc4576b938e37cb86"),
        },
        {
           fromHex("a746273228122f381c3b46e4f1"),
           fromHex("ec715d44cf1465caae6e0d620cd4aa745d7240ac5fb7a18a8bf84b5ae27f411db289313dfbc5396fe40ee2789257c56f"),
        },
        {
           fromHex("d8dc8fdefbdce9d44e4cbafe78447bae3b5436102a"),
           fromHex("fba3df2b335ccbf2ec21be28c4d818e54e8ae3f1d0891df51133f9b45a2c40d0b82236798530de21d0119fc45f6300df"),
        },
        {
           fromHex("aecbb02759f7433d6fcb06963c74061cd83b5b3ffa6f13c6"),
           fromHex("393fa754df52f0a8a9ac1ad465360b272374e68db174b26e0bef195ecab4eff42d0d7ca0ad8adfb3f2d408bf6be13dee"),
        },
        {
           fromHex("7bc84867f6f9e9fdc3e1046cae3a52c77ed485860ee260e30b15"),
           fromHex("f9ec1e70ff45da6c47635387c36bddb9ee497ee5b65e62ff99ce47627d6331f7156e2436d53ea4efb3037eeaafd95e49"),
        },
        {
           fromHex("a963c3e895ff5a0be4824400518d81412f875fa50521e26e85eac90c04"),
           fromHex("5902e8d77c7b96cc0cce8b2a8de2c69c813701aff7ed048eeb137babd1a76cf2646fed00129d7f2f495ad0652eedf8cb"),
        },
        {
           fromHex("03a18688b10cc0edf83adf0a84808a9718383c4070c6c4f295098699ac2c"),
           fromHex("99272657ea32e389011686b8d1c515d9612d869b8519f2485f62de7c30c8c5d702bab43b73194a12ea144337d515a3c3"),
        },
        {
           fromHex("9f2fcc7c90de090d6b87cd7e9718c1ea6cb21118fc2d5de9f97e5db6ac1e9c10"),
           fromHex("93f1da7d842b550d1bfa8debf8ee4f595a0a3b056c141b202e025d890e4ae1b310fd0874e060b33be865d2e163938388"),
        },
    }

    h := New384()

    for i, test := range tests {
        h.Reset()
        h.Write(test.msg)
        sum := h.Sum(nil)

        if !bytes.Equal(sum, test.md) {
            t.Errorf("[%d] New384 fail, got %x, want %x", i, sum, test.md)
        }

        // =====

        sum2 := Sum384(test.msg)

        if !bytes.Equal(sum2[:], test.md) {
            t.Errorf("[%d] Sum384 fail, got %x, want %x", i, sum2, test.md)
        }
    }
}

func Test_Hash512_Check(t *testing.T) {
   tests := []testData{
        {
           fromHex(""),
           fromHex("5cd7436a91e27fc809d7015c3407540633dab391127113ce6ba360f0c1e35f404510834a551610d6e871e75651ea381a8ba628af1dcf2b2be13af2eb6247290f"),
        },
        {
           fromHex("cc"),
           fromHex("7da1be62a813a8e24d200671cffb1d0be79d2bc176ff0b163b11eded2414ef66261ff52c745383442bc7f1884d5166f26f41d335fc2d2fdb2f93b24b8d079265"),
        },
        {
           fromHex("41fb"),
           fromHex("3253d2db0d57862d6deec1033f27e373d3becbab7fa74c9b3ec1d041bbca8978c19e34e3e726a7c163c7d6a996897a5db80b21b385c47e8e3a3aee6023388cf2"),
        },
        {
           fromHex("1F877C"),
           fromHex("af015a97b6996ed048f32b3a6c209e6a2daeacd4f61eb62eaa31c68328ee5790b0681245ebe1ecec4c0dd7f9008672d28a0424406998ec02518f023b3c27dcde"),
        },
        {
           fromHex("c1ecfdfc"),
           fromHex("4bdac806bd3111b72e91df166102ae846e44f6b9cdfe0aaac07dd9730d4bebeb0860919887518db8d1f32e32c72efc35bbe487899cc5fe3388caa8ce096975e2"),
        },
        {
           fromHex("1f66ab4185ed9b6375"),
           fromHex("20055d024bbe68c70e5580cb904bc9d6d867f430fd630e5684c1d5178ef16336b067f5a15c7567a6465b6e95d4bf027b4ee1add9bcf62a5c1169e183afc2344e"),
        },
        {
           fromHex("a746273228122f381c3b46e4f1"),
           fromHex("890bf4df9afce7b9f400af26842980c2e97d3dea6d7ef196623c45a5ea94f3598999f6a988941343d2528f0e732d1f483597a2f71859f17f821877be6320852b"),
        },
        {
           fromHex("d8dc8fdefbdce9d44e4cbafe78447bae3b5436102a"),
           fromHex("317c7c79e6c634ecc7541b590aced974c785e4ff2012caffe3dde9584f60f802297cf242cae9a5dd1e13324c9f104af319ea7e56795bd8060be9408e465e71c9"),
        },
        {
           fromHex("aecbb02759f7433d6fcb06963c74061cd83b5b3ffa6f13c6"),
           fromHex("fe2cfe47de1970a2cf768243fb8c312d93035edfb361de232e4f947f825cf670a83686b12da1569da9d69f7415bd74f39cdbb418eaf160af0335e3af7a5e3eb2"),
        },
        {
           fromHex("7bc84867f6f9e9fdc3e1046cae3a52c77ed485860ee260e30b15"),
           fromHex("50436276410bdcb25cedf0319e0093fed4fb6f5bbc3a279f305ee94e0e52c78624c8d0e9346e52baa8325a46f63430bec92606b4964dac7189a26b3e214e2c63"),
        },
        {
           fromHex("a963c3e895ff5a0be4824400518d81412f875fa50521e26e85eac90c04"),
           fromHex("2c67bc87b9da7fb42767a128e6b1d2cab04a057d0f179617f483e8a387f5f67f6b64664f7400f2c7b2120ebf7c228347ca5a68d4c7d2a7d7a9d26eddf2364a29"),
        },
        {
           fromHex("03a18688b10cc0edf83adf0a84808a9718383c4070c6c4f295098699ac2c"),
           fromHex("39fe4f590a6093e5fbbb59a54c9150a2ff944a921938b9a4c97c599d45d78255274456bb5ec73676b610a91270d466d2e4079a799da9d7057c015ce9bed1fe71"),
        },
        {
           fromHex("84fb51b517df6c5accb5d022f8f28da09b10232d42320ffc32dbecc3835b29"),
           fromHex("9940753170edd23210a8ddd74c70170193f34231e9ca03ffeadaba0e15ea0d4be1772044f0b65e734adf1602730487395c104e6e5e9f9dea1b9359bffd264e76"),
        },
        {
           fromHex("9f2fcc7c90de090d6b87cd7e9718c1ea6cb21118fc2d5de9f97e5db6ac1e9c10"),
           fromHex("011ca7f5da5b73869f6e30139002b89cfaf3c54d825ec002ab205ed8b2a317e1037b75f4bcad6f0d7dd2462fee5924078351eba55313baf02301fe451920acd4"),
        },
    }

    h := New512()

    for i, test := range tests {
        h.Reset()
        h.Write(test.msg)
        sum := h.Sum(nil)

        if !bytes.Equal(sum, test.md) {
            t.Errorf("[%d] New512 fail, got %x, want %x", i, sum, test.md)
        }

        // =====

        sum2 := Sum512(test.msg)

        if !bytes.Equal(sum2[:], test.md) {
            t.Errorf("[%d] Sum512 fail, got %x, want %x", i, sum2, test.md)
        }
    }
}
