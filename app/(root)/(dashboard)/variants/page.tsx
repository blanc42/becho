"use client"
import { Button } from "@/components/ui/button";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import Link from "next/link";

export default function VariantsPage() {

  return (
    <>
      <div  className="flex items-baseline justify-between">
        <h1 className="text-3xl font-SemiBold">Variants</h1>
        <Button>
          <Link href="/variants/add">
            Add Variant
          </Link>
        </Button>
      </div>
      <div>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>
                <TableCell>Name</TableCell>
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>Name</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </div>
    </>
  )
}