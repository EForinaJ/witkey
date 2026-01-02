declare namespace Account {
    namespace Params {
        interface Model {
            name: string | null;
            email: string | null;
            address: string[];
            sex: number | null;
            avatar: string | null;
            description: string | null;
            tags: string[];
        }
        interface ChangePass {
            oldPass: string | null;
            newPass: string | null;
            confirmPass: string | null;
        }
    }
    namespace Response {
        /** 登录参数 */
        type Info = {
            // permission: string[]
            id: number
            name: string| null;
            avatar?: string| null;
            email: string | null;
            address: string[];
            sex: number | null;
            witkey: {
                id: number
                name: string| null;
                title: string| null;
                game: string| null;
                commission: number| null;
                album: string[];
                rate: number| null;
            };
            loginIp?: string| null;
        }
    }
}