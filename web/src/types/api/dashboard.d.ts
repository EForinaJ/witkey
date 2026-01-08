declare namespace Dashboard {
    namespace Params {
    }
    namespace Response {
       
        type Info = {
            id: number;
            code: string;
            product: {
                id: number;
                name: string;
                pic: string;
                game: string;
                category: string;
            };
            quantity: number;
            unit: string;
            type: number;
            totalAmount: number;
            status: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            comparedToYesterday:number;
            TodayCommission:number;
            distributeList: {
                id: number;
                code: string;
                product: {
                    id: number;
                    name: string;
                    pic: string;
                    game: string;
                    category: string;
                };
                quantity: number;
                unit: string;
                type: number;
                totalAmount: number;
                status: number;
                createTime: string;
            }[];
        }
    }
}