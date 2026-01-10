declare namespace Account {
    namespace Params {
        interface Model {
            name: string | null;
            address: string[];
            sex: number | null;
            avatar: string | null;
            description: string | null;
            birthday: number | null;
            album: string[];
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
            birthday: number | null;
            address: string[];
            sex: number | null;
            title: string| null;
            game: string| null;
            commission: number| null;
            album: string[];
            rate: number| null;
            description: string| null;
            loginIp?: string| null;
        }
    }
}