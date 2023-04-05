import { createEffect, createSignal } from "solid-js";
import { useParams, useNavigate } from "@solidjs/router";
import { supabase } from "../../supabase-client";
import CardWrapper from "../../components/CardWrapper/CardWrapper";
// @ts-expect-error
import styles from "./Details.module.css";
import { Transaction, TokenTransfer } from "../../schema/supabase-types";

const TransactionsDetails = () => {
  const [details, setDetails] = createSignal<Transaction | any>({});
  const [transfers, setTransfers] = createSignal<TokenTransfer[]>([]);
  const params = useParams();

  const detailsTableProps = {
    onClick: {
      action: () => {},
      prop: "id",
    },
    schema: [
      { display: "transfer id", objProp: ["id"] },
      { display: "sender/recipient", objProp: ["sender", "recipient"] },
      { display: "value", objProp: ["value"] },
    ],
  };

  const fetchDetails = async () => {
    try {
      const { data, error } = await supabase
        .from("transaction")
        .select("*")
        .match({ id: params.id });
      if (data?.[0]) {
        setDetails(data[0]);
      } else {
        throw new Error("Failed to load transaction details");
      }
    } catch (error) {
      console.error(error.message);
    }
  };

  const fetchTransfers = async () => {
    try {
      const { data, error } = await supabase
        .from("token_transfer")
        .select("*")
        .match({ transaction_id: params.id });
      if (data) {
        setTransfers(data);
      } else {
        throw new Error("Failed to load transfers");
      }
    } catch (error) {
      console.error(error.message);
    }
  };

  createEffect(async () => {
    await fetchDetails();
    await fetchTransfers();
  }, []);

  return (
    <div class={styles.details}>
      <div class={styles["details-card"]}>
        <h1> Details of Transaction</h1>
        <span>
          {" "}
          <b>Hash:</b> {details().hash}{" "}
        </span>
        <span>
          <b>Block: </b>
          {details().block_number}
        </span>
        <span>
          <b>From: </b>
          {details().from}
        </span>
        <span>
          <b>To: </b>
          {details().to}{" "}
        </span>
        <span>
          {" "}
          <b>Value: </b>&nbsp;&nbsp;{details().eth_value}
        </span>
      </div>
      <CardWrapper
        columns={detailsTableProps}
        data={transfers()}
        title={"Transfer list"}
        display={"small"}
        filters={["id", "sender", "recipient"]}
        trimmed={false}
      />
    </div>
  );
};

export default TransactionsDetails;
