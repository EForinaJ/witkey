declare namespace Withdraw {
    namespace Params {
        interface Query {
            page: number;
            limit: number;
            code?: string;
            status?: number;
        }

    }
    namespace Response {
       
        type Info = {
            id: number;
            code: string;
            amount: number;
            settledAmount: number;
            serviceFee: number;
            type: number;
            status: number;
            createTime: string;
        }
        type List = Api.Common.PaginatedResponse<Info>

        type Detail = {
            id: number;
            code: string;
            amount: number;
            settledAmount: number;
            serviceFee: number;
            type: number;
            name: string;
            number: string;
            status: number;
            reason: string;
            createTime: string;
        }
    }
}